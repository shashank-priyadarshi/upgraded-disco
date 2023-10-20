package redis

type RedisDatabase struct{}

func (rd *RedisDatabase) Create(config, data interface{}) error{return nil}
func(rd *RedisDatabase) Get(config, data interface{}) (interface{}, error){return nil,nil}
func(rd *RedisDatabase) Update(config, data interface{}) error{return nil}
func(rd *RedisDatabase) Delete(config, data interface{}) error{return nil}