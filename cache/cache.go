package cache

import (
	"go.uber.org/zap"
)

type Cacher interface {
	Get(logger *zap.Logger, key string) (string, error)
	Set(logger *zap.Logger, key, value string, expireInSec int) error
	Delete(logger *zap.Logger, key string) error
}