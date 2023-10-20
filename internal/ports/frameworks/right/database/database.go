package database

type Database interface{
	Create(config, data interface{}) error
	Get(config, data interface{}) (interface{}, error)
	Update(config, data interface{}) error
	Delete(config, data interface{}) error
}