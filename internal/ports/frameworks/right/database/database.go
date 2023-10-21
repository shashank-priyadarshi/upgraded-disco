package database

type Database interface {
	Create(data interface{}) error
	Get(data interface{}) (interface{}, error)
	Update(data interface{}) error
	Delete(data interface{}) error
}
