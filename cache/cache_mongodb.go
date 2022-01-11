package cache

import (
	"fmt"
	apierrors "gitlab.globoi.com/globoid/globoid-kids/api_errors"
	"time"

	"github.com/go-redis/redis"
	"gitlab.globoi.com/globoid/globoid-kids/context"
	"go.uber.org/zap"
)

type mongoCache struct {
	client *.Client
}

func NewMongoCache(database int) (*mongoCache, error) {
	logger, err := context.GetSimpleLogger()
	apiConfig := context.GetAPIConfig()
	env := context.GetAPIConfig().AppEnv

	var redisClient *redis.Client

	if env == "prod" { // use sentinel redis
		redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			ReadTimeout:   time.Duration(1) * time.Second,
			WriteTimeout:  time.Duration(1) * time.Second,
			MasterName:    apiConfig.SentinelConfig.MasterName,
			SentinelAddrs: apiConfig.SentinelConfig.SentinelAddrs,
			Password:      apiConfig.SentinelConfig.Password,
			DB:            database,
			PoolSize:      5000,
		})
	} else { // use redis
		redisClient = redis.NewClient(&redis.Options{
			ReadTimeout:  time.Duration(1) * time.Second,
			WriteTimeout: time.Duration(1) * time.Second,
			Addr:         apiConfig.RedisConfig.RedisURL,
			Password:     apiConfig.RedisConfig.RedisPassword,
			DB:           database,
			PoolSize:     5000,
		})
	}

	if _, err = redisClient.Ping().Result(); err != nil {
		logger.Error("Error when pinging on created redis client", zap.Error(err))
	}

	return &RedisCache{
		client: redisClient,
	}, nil
}

func (r RedisCache) Get(logger *zap.Logger, key string) (string, error) {
	cacheValue, err := r.client.Get(key).Result()

	if err == redis.Nil {
		logger.Debug(fmt.Sprintf("Key %s not found on cache", key))
		return "", apierrors.ErrorNotFoundOnCache
	}

	if err != nil {
		logger.Error("Error when getting value from redis cache", zap.Error(err))
		return "", apierrors.ErrorGetCacheValue
	}

	logger.Info("Key entry found")
	return cacheValue, nil
}

func (r RedisCache) Set(logger *zap.Logger, key, value string, expireInSec int) error {
	expire := time.Duration(expireInSec) * time.Second

	err := r.client.Set(key, value, expire).Err()
	if err != nil {
		logger.Error("Error when setting value on redis cache", zap.Error(err))
	}

	logger.Info(fmt.Sprintf("Set key %s on cache with expire in %d seconds", key, expireInSec))
	return err
}

func (r RedisCache) Delete(logger *zap.Logger, key string) error {
	err := r.client.Del(key).Err()
	if err == redis.Nil {
		logger.Debug(fmt.Sprintf("Key %s not found on cache", key))
		return apierrors.ErrorNotFoundOnCache
	}
	if err != nil {
		logger.Error("Error when deletting key from redis cache", zap.Error(err))
		return err
	}

	logger.Info("Key entry deleted")
	return nil
}
