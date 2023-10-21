package mongodb

import (
	"context"
	"fmt"
	databases "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

type MongoDatabase struct {
	client *mongo.Client
	logger zap.Logger
}

func NewMongoDBInstance(log, config interface{}) (*MongoDatabase, error) {
	if config == nil {
		return &MongoDatabase{}, fmt.Errorf("MongoDB config cannot be nil")
	}
	cnf := config.(databases.MongoDBConfig)
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
		logger: log.(zap.Logger),
	}, nil
}

func (rd *MongoDatabase) Create(data interface{}) error             { return nil }
func (rd *MongoDatabase) Get(data interface{}) (interface{}, error) { return nil, nil }
func (rd *MongoDatabase) Update(data interface{}) error             { return nil }
func (rd *MongoDatabase) Delete(data interface{}) error             { return nil }
