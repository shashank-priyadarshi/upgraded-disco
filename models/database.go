package models

type Fields = map[string]string

type DBConfig struct {
	Username, Password, Host               string
	Database                               interface{}
	MaxIdleConnections, MaxOpenConnections int
}

type RedisConfig struct {
	DBConfig
}

type MongoDBConfig struct {
	DBConfig
}

type MariaDBConfig struct {
	DBConfig
}

type RedisPayload struct {
	Key   string
	Value interface{}
}

type MongoDBPayload struct {
	Database, Collection string
	Data                 interface{}
}

type MariaDBPayload struct {
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50);unique_index"`
	Password string `gorm:"type:varchar(50)"`
	Username string `gorm:"type:varchar(10);unique_index"`
}
