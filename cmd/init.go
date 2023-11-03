package main

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/app/config"
	models "github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/core/domain"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/plugins"
	"go.uber.org/zap"
	"os"
)

func init() {
	logger := zap.Logger{} // TODO: Set log levels, log mode and log source

	configSource := config.Source{
		Source:     os.Getenv("CONFIG_SOURCE"),
		SourcePath: os.Getenv("CONFIG_SOURCE_PATH"),
	}
	appConfig, err := configSource.Load()
	if err != nil {
		panic(err)
	}

	maria, err := databases.NewDatabase("Maria", logger, appConfig.DBConfig)
	if err !=
		nil {
		panic(err)
	}
	mongo, err := databases.NewDatabase("Mongo", logger, appConfig.DBConfig)
	if err != nil {
		panic(err)
	}
	redis, err := databases.NewDatabase("Redis", logger, appConfig.DBConfig)
	if err != nil {
		panic(err)
	}

	pluginConf := plugins.Plugin{
		Databases: models.Databases{
			Redis: redis,
			Maria: maria,
			Mongo: mongo,
		},
		Log: logger,
	}
	initPlugin(&pluginConf)

	// init server
	// accepts config, maria, redis, mongo and logger
}

func initPlugin(pluginConf *plugins.Plugin) {
	// Get OS: set work directory, initialise directory for built binaries
	// Read supported languages, read commands, check if dependencies available
	// asynchronously iterate over plugin source directory and build all plugins
}
