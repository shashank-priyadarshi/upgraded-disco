package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	databases "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/core"
	"go.uber.org/zap"
)

type RedisDatabase struct {
	client *redis.Client
	logger zap.Logger
}

func NewRedisInstance(log, config interface{}) (*RedisDatabase, error) {
	if config == nil {
		return &RedisDatabase{}, fmt.Errorf("redis config cannot be nil")
	}
	cnf := config.(databases.RedisConfig)
	rDBClient := redis.NewClient(&redis.Options{
		Addr:           cnf.Host,
		Username:       cnf.Username,
		Password:       cnf.Password,
		DB:             cnf.Database.(int),
		MaxIdleConns:   cnf.MaxIdleConnections,
		MaxActiveConns: cnf.MaxOpenConnections,
		TLSConfig:      nil,
	})
	if err := rDBClient.Ping(context.Background()); err != nil {
		return &RedisDatabase{}, fmt.Errorf("error initilizing Redis DB: %v", err)
	}
	return &RedisDatabase{
		client: rDBClient,
		logger: log.(zap.Logger),
	}, nil
}

func (rd *RedisDatabase) Create(config, data interface{}) error {
	return nil
}

func (rd *RedisDatabase) Get(config, data interface{}) (interface{}, error) { return nil, nil }

func (rd *RedisDatabase) Update(config, data interface{}) error { return nil }

func (rd *RedisDatabase) Delete(config, data interface{}) error { return nil }
