package main

import (
	"fmt"
	driver "github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/fasthttp"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/plugins"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/config"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {
	appLogger := logger.NewLogger("INFO", "")
	appLogger.Infof("Beginning application initialisation")
	configSource := config.Source{
		ConfigSource: os.Getenv("CONFIG_SOURCE"),
		ConfigPath:   os.Getenv("CONFIG_PATH"),
	}
	appConfig, err := configSource.Load()
	if err != nil {
		panic(fmt.Sprintf("error initialising config from %s: %v", configSource.ConfigSource, err))
	}
	appLogger.Infof("Loaded application config")

	app := application.NewApplication(appLogger, appConfig.DBConfig)
	appLogger.Infof("Initialised services")
	router := driver.NewRouter(&appConfig).SetRouter(app, appLogger)
	appLogger.Infof("Setup up application routers")
	appLogger.Infof("Starting up application server on port %v", appConfig.ServerConfig.Port)
	serverErr := fasthttp.ListenAndServe(appConfig.ServerConfig.Port, router.Handler)
	appLogger.Infof("error while serving: %v", serverErr)
}

// TODO
func initPlugin(pluginConf *plugins.Plugin) {
	// Get OS: set work directory, initialise directory for built binaries
	// Read supported languages, read commands, check if dependencies available
	// asynchronously iterate over plugin source directory and build all plugins
}
