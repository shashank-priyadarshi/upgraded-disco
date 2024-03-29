package models

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Config struct {
	ServerConfig
	DBConfig map[string]DBConfig
	PluginsConfig
}

type Repository struct {
	Log                     logger.Logger
	Cache, MariaDB, MongoDB ports.Database
}
