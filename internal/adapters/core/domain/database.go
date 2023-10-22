package models

type Fields = map[string]string

type DBConfig struct {
	Username, Password, Host               string
	Database                               interface{}
	MaxIdleConnections, MaxOpenConnections int
}

// QueryField RedisPayload MongoDBPayload MariaDBPayload TODO: Use value to QueryField.Key to get the value of respective struct property from DB payloads
type QueryField struct {
	Fields
}

type UpdateField struct {
	Fields
}

type RedisConfig struct {
	DBConfig
}

type RedisPayload struct {
	Key   string
	Value interface{}
}

type MongoDBConfig struct {
	DBConfig
}

type MongoDBPayload struct {
	Database, Collection string
	Data                 interface{}
}

type MariaDBConfig struct {
	DBConfig
}

type MariaDBPayload struct {
	Name     string `gorm:"type:varchar(50)"`
	Email    string `gorm:"type:varchar(50);unique_index"`
	Password string `gorm:"type:varchar(50)"`
	Username string `gorm:"type:varchar(10);unique_index"`
}
