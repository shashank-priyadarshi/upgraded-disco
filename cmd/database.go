package main

import (
	r "github.com/redis/go-redis/v9"
	databases "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/core"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func redis(log zap.Logger, redisConfig databases.RedisConfig) (*r.Client, error) {
	return &r.Client{}, nil
}

func mongodb(log zap.Logger, mongodbConfig databases.MongoDBConfig) (*mongo.Client, error) {
	return &mongo.Client{}, nil
}

func mariadb(log zap.Logger, mariadbConfig databases.MariaDBConfig) (*gorm.DB, error) {
	return &gorm.DB{}, nil
}
