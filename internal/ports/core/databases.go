package databases

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
