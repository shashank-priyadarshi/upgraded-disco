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

func (rd *RedisDatabase) Create(data interface{}) error {
	payload := data.(databases.RedisPayload)
	if err := rd.client.Set(context.Background(), payload.Key, payload.Value, 0); err != nil {
		return fmt.Errorf("error putting value for Key %s to redis cache: %s", payload.Key, err)
	}
	return nil
}

func (rd *RedisDatabase) Get(data interface{}) (interface{}, error) {
	payload := data.(databases.RedisPayload)
	value := rd.client.Get(context.Background(), payload.Key)
	if value.Val() == redis.Nil.Error() {
		return nil, fmt.Errorf("key %s does not exist in redis cache", payload.Key)
	}
	return value.Val(), nil
}

func (rd *RedisDatabase) Update(data interface{}) error {
	payload := data.(databases.RedisPayload)
	if err := rd.client.Set(context.Background(), payload.Key, payload.Value, 0); err != nil {
		return fmt.Errorf("error updating value for key %s in redis cache: %s", payload.Key, err)
	}
	return nil
}

func (rd *RedisDatabase) Delete(data interface{}) error {
	payload := data.(databases.RedisPayload)
	if err := rd.client.Del(context.Background(), payload.Key); err != nil {
		return fmt.Errorf("error removing key %s from redis cache: %s", payload.Key, err)
	}
	return nil
}
