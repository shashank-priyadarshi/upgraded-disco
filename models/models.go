package models

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type Config struct {
	ServerConfig
	DBConfig
	PluginsConfig
}

type Databases struct {
	Maria, Redis, Mongo ports.Database
}
