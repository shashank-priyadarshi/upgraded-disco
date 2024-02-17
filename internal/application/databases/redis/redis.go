package redis

import (
	"context"
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/databases/batch_ops"
	"strconv"

	"github.com/redis/go-redis/v9"
	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type RedisDatabase struct {
	client   *redis.Client
	logger   logger.Logger
	BatchOps batch_ops.BatchOps
}

func NewRedisInstance(log logger.Logger, config interface{}) (*RedisDatabase, error) {
	if config == nil {
		return &RedisDatabase{}, fmt.Errorf("redis config cannot be nil")
	}
	cnf := config.(models.DBConfig)
	database, err := strconv.Atoi(cnf.Database[0].(string))
	if err != nil {
		return &RedisDatabase{}, fmt.Errorf("error parsing redis database name %v: %v", cnf.Database, err)
	}
	rDBClient := redis.NewClient(&redis.Options{
		Addr:           cnf.Host,
		Username:       cnf.Username,
		Password:       cnf.Password,
		DB:             database,
		MaxIdleConns:   cnf.MaxIdleConnections,
		MaxActiveConns: cnf.MaxOpenConnections,
		TLSConfig:      nil,
	})
	//if err := rDBClient.Ping(context.Background()); err != nil {
	//	return &RedisDatabase{}, fmt.Errorf("error initilizing Redis DB: %v", err)
	//}

	rDB := &RedisDatabase{
		client: rDBClient,
		logger: log,
	}
	rDB.BatchOps = batch_ops.BatchOps{rDB}

	return rDB, nil
}

func (rd *RedisDatabase) Exists(data interface{}) bool {
	if data == nil {
		return false
	}
	// TODO
	return true
}

func (rd *RedisDatabase) Create(data interface{}) (interface{}, error) {
	payload := data.(models.RedisPayload)
	if err := rd.client.Set(context.Background(), payload.Key, payload.Value, 0); err != nil {
		return nil, fmt.Errorf("error putting value for Key %s to redis cache: %s", payload.Key, err)
	}
	return nil, nil
}

func (rd *RedisDatabase) Get(data interface{}) (interface{}, error) {
	payload := data.(models.RedisPayload)
	value := rd.client.Get(context.Background(), payload.Key)
	if value.Val() == redis.Nil.Error() {
		return nil, fmt.Errorf("key %s does not exist in redis cache", payload.Key)
	}
	return value.Val(), nil
}

func (rd *RedisDatabase) Update(data interface{}) (interface{}, error) {
	payload := data.(models.RedisPayload)
	if err := rd.client.Set(context.Background(), payload.Key, payload.Value, 0); err != nil {
		return nil, fmt.Errorf("error updating value for key %s in redis cache: %s", payload.Key, err)
	}
	return nil, nil
}

func (rd *RedisDatabase) Delete(data interface{}) (interface{}, error) {
	payload := data.(models.RedisPayload)
	if err := rd.client.Del(context.Background(), payload.Key); err != nil {
		return nil, fmt.Errorf("error removing key %s from redis cache: %s", payload.Key, err)
	}
	return nil, nil
}
