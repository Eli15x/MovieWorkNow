package storage

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"time"
	"fmt"
	"github.com/Eli15x/MovieWorkNow/utils"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/labstack/gommon/log"
)

var (
	once          sync.Once
	mongoInstance MongoDB
)

type MongoDB interface {
	Insert(ctx echo.Context, collName string, doc interface{}) (interface{}, error)
	Find(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) error
	FindOne(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) error
	Count(ctx echo.Context, collName string, query map[string]interface{}) (int64, error)
	UpdateOne(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) (*mongo.UpdateResult, error)
	Remove(ctx echo.Context, collName string, query map[string]interface{}) error
	WithTransaction(ctx echo.Context, fn func(context.Context) error) error
	Initialize(ctx context.Context) error
	Disconnect()
}

type mongodbImpl struct {
	client *mongo.Client
	dbName string
}

func GetInstance() MongoDB {
	once.Do(func() {
		mongoInstance = &mongodbImpl{}
	})
	return mongoInstance
}

func (m *mongodbImpl) Initialize(ctx context.Context) error {

	clientOptions := options.Client().
	ApplyURI("mongodb+srv://elisacds:elisacds@cluster0.e7uxp.mongodb.net/MovieWorkNow?retryWrites=true&w=majority")

	client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }

	//Está dando proble	

	/*err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}*/

	m.dbName = "MovieWorkNow"
	m.client = client
	return nil
}

func (m *mongodbImpl) WithTransaction(ctx echo.Context, fn func(context.Context) error) error {
	return m.client.UseSession(ctx.Request().Context(), func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}
		err = fn(sessionContext)
		if err != nil {
			return sessionContext.AbortTransaction(sessionContext)
		}
		return sessionContext.CommitTransaction(sessionContext)
	})
}

// Insert stores documents in the collection
func (m *mongodbImpl) Insert(ctx echo.Context, collName string, doc interface{}) (interface{}, error) {

	insertedObject, err := m.client.Database(m.dbName).Collection(collName).InsertOne(ctx.Request().Context(), doc)
	if insertedObject == nil {
		fmt.Println(err)
		return nil, err
	}
	return insertedObject.InsertedID, err
}

// Find finds all documents in the collection
func (m *mongodbImpl) Find(echoCtx echo.Context, collName string, query map[string]interface{}, doc interface{}) error {
	ctx := echoCtx.Request().Context()
	segment := utils.StartSegmentWithDatastoreProduct(echoCtx, "Mongo.Find", newrelic.DatastoreMongoDB, "Find", collName)
	defer segment.End()

	cur, err := m.client.Database(m.dbName).Collection(collName).Find(ctx, query)
	if err != nil {
		return err
	}

	resultv := reflect.ValueOf(doc)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return errors.New("failed to return array response")
	}

	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elem := slicev.Type().Elem()

	i := 0
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		elemp := reflect.New(elem)
		err := cur.Decode(elemp.Interface())
		if err != nil {
			return err
		}
		slicev = reflect.Append(slicev, elemp.Elem())
		slicev = slicev.Slice(0, slicev.Cap())
		i++
	}

	resultv.Elem().Set(slicev.Slice(0, i))
	return nil
}

// FindOne finds one document in mongo
func (m *mongodbImpl) FindOne(ctx echo.Context, collName string, query map[string]interface{}, doc interface{}) error {
	segment := utils.StartSegmentWithDatastoreProduct(ctx, "Mongo.FindOne", newrelic.DatastoreMongoDB, "FindOne", collName)
	defer segment.End()

	return m.client.Database(m.dbName).Collection(collName).FindOne(ctx.Request().Context(), query).Decode(doc)
}

// UpdateOne updates one or more documents in the collection
func (m *mongodbImpl) UpdateOne(ctx echo.Context, collName string, selector map[string]interface{}, update interface{}) (*mongo.UpdateResult, error) {

	segment := utils.StartSegmentWithDatastoreProduct(ctx, "Mongo.UpdateOne", newrelic.DatastoreMongoDB, "UpdateOne", collName)
	defer segment.End()

	updateResult, err := m.client.Database(m.dbName).Collection(collName).UpdateOne(ctx.Request().Context(), selector,update)
	fmt.Println(err)
	return updateResult, err
}

// Remove one or more documents in the collection
func (m *mongodbImpl) Remove(ctx echo.Context, collName string, selector map[string]interface{}) error {
	segment := utils.StartSegmentWithDatastoreProduct(ctx, "Mongo.Remove", newrelic.DatastoreMongoDB, "Remove", collName)
	defer segment.End()

	_, err := m.client.Database(m.dbName).Collection(collName).DeleteOne(ctx.Request().Context(), selector)
	return err
}

// Count returns the number of documents of the query
func (m *mongodbImpl) Count(ctx echo.Context, collName string, query map[string]interface{}) (int64, error) {
	segment := utils.StartSegmentWithDatastoreProduct(ctx, "Mongo.Count", newrelic.DatastoreMongoDB, "Count", collName)
	defer segment.End()

	return m.client.Database(m.dbName).Collection(collName).CountDocuments(ctx.Request().Context(), query)
}

func (m *mongodbImpl) Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	_ = m.client.Disconnect(ctx)
}
