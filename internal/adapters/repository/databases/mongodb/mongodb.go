package mongodb

import (
	"context"
	"fmt"

	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDatabase struct {
	client *mongo.Client
	logger logger.Logger
}

func NewMongoDBInstance(log logger.Logger, config interface{}) (*MongoDatabase, error) {
	if config == nil {
		return &MongoDatabase{}, fmt.Errorf("MongoDB config cannot be nil")
	}
	cnf := config.(models.DBConfig)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/", cnf.Username, cnf.Password, cnf.Host)))
	if err != nil {
		return &MongoDatabase{}, fmt.Errorf("error initializing mongo db client: %v", err)
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return &MongoDatabase{}, fmt.Errorf("error connecting to mongo db: %v", err)
	}
	return &MongoDatabase{
		client: client,
		logger: log,
	}, nil
}

func (rd *MongoDatabase) Create(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}
	var insertResult *mongo.InsertOneResult
	var err error
	payload := data.(models.MongoDBPayload)
	if payload.Data == nil {
		return nil, fmt.Errorf("database create query cannot be nil")
	}
	if insertResult, err = rd.client.Database(payload.Database).Collection(payload.Collection).InsertOne(context.Background(), payload.Data); err != nil {
		return nil, fmt.Errorf("error creating new document %+v in database %s and collection %s: %s", payload.Data, payload.Database, payload.Collection, err)
	}
	return insertResult, nil
}

func (rd *MongoDatabase) Get(data interface{}) (interface{}, error) {
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}
	var results interface{}
	var cursor *mongo.Cursor
	var err error
	payload := data.(models.MongoDBPayload)
	if payload.Data == nil {
		return nil, fmt.Errorf("database get query cannot be nil")
	}
	filter := payload.Data.(models.Fields)
	if cursor, err = rd.client.Database(payload.Database).Collection(payload.Collection).Find(context.Background(), filter); err != nil {
		return nil, fmt.Errorf("error finding document %+v in database %s and collection %s: %s", payload.Data, payload.Database, payload.Collection, err)
	}
	defer cursor.Close(context.Background())
	return cursor.All(context.Background(), &results), nil
}

func (rd *MongoDatabase) Update(data interface{}) (interface{}, error) {
	var updateResult *mongo.UpdateResult
	var err error
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}
	payload := data.(models.MongoDBPayload)
	if payload.Data == nil {
		return nil, fmt.Errorf("database update queries cannot be nil")
	}
	queries := payload.Data.([]models.Fields)
	filter, update := queries[0], queries[1]
	if updateResult, err = rd.client.Database(payload.Database).Collection(payload.Collection).UpdateOne(context.Background(), filter, update); err != nil {
		return nil, fmt.Errorf("error updating document %+v in the database %s and collection %s: %s", payload.Data, payload.Database, payload.Collection, err)
	}
	return struct {
		MatchedCount  int64
		UpsertedCount int64
		ModifiedCount int64
	}{
		MatchedCount:  updateResult.MatchedCount,
		UpsertedCount: updateResult.UpsertedCount,
		ModifiedCount: updateResult.ModifiedCount,
	}, nil
}

func (rd *MongoDatabase) Delete(data interface{}) (interface{}, error) {
	var deleteResult *mongo.DeleteResult
	var err error
	if data == nil {
		return nil, fmt.Errorf("input data cannot be nil")
	}
	payload := data.(models.MongoDBPayload)
	if payload.Data == nil {
		return nil, fmt.Errorf("database delete query cannot be nil")
	}
	filter := payload.Data.(models.Fields)
	if deleteResult, err = rd.client.Database(payload.Database).Collection(payload.Collection).DeleteOne(context.Background(), filter); err != nil {
		return nil, fmt.Errorf("error deleting document %+v in database %s and collection %s: %s", payload.Data, payload.Database, payload.Collection, err)
	}
	return deleteResult.DeletedCount, nil
}
