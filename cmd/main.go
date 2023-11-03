package main

import (
	"os"

	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/config"
	models "github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/plugins"
	"go.uber.org/zap"
)

func main() {
	//	TODO: Add Strategy Pattern to add services as required and start server
	//	TODO: Structure init for all services
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
