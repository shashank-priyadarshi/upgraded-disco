package models

import "github.com/shashank-priyadarshi/upgraded-disco/internal/ports/frameworks/right/database"

type Config struct {
	ServerConfig
	DBConfig
	PluginsConfig
}

type Databases struct {
	Maria, Redis, Mongo database.Database
}
