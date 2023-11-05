package plugin

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type pluginRepo struct {
	ports.Database
}

func NewPluginRepo(log logger.Logger, config interface{}) ports.PluginRepo {
	db, err := databases.NewDatabase("", log, config)
	if err != nil {
		log.Errorf("failed to initialise schedule repo: %v", err)
		return nil
	}
	return &pluginRepo{
		db,
	}
}
