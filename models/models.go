package models

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	logger "github.com/shashank-priyadarshi/utilities/logger/ports"
)

type Config struct {
	ServerConfig
	DBConfig map[string]DBConfig
	PluginsConfig
}

type Repository struct {
	Log                     logger.Logger
	Cache, MariaDB, MongoDB ports.Repository
}
