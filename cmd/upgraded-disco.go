package main

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/config"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application"
	driver "github.com/shashank-priyadarshi/upgraded-disco/internal/application/fasthttp"
	"github.com/shashank-priyadarshi/utilities/logger"
	"github.com/valyala/fasthttp"
	"os"
)

func main() {

	appLogger, err := logger.NewLogger("slog", "Info", "text", true)
	if err != nil {
		panic(fmt.Sprintf("error initialising default logger: %v", err))
	}
	appLogger.Info("Beginning application initialisation")

	configSource := config.Source{
		ConfigSource: os.Getenv("CONFIG_SOURCE"),
		ConfigPath:   os.Getenv("CONFIG_PATH"),
	}

	appConfig, err := configSource.Load(appLogger)
	if err != nil {
		panic(fmt.Sprintf("error initializing config from %s: %v", configSource.ConfigSource, err))
	}
	appLogger.Info("Loaded application config")

	app := application.NewApplication(appLogger, appConfig.DBConfig, &appConfig.PluginsConfig)
	appLogger.Info("Initialised services")

	router := driver.NewRouter(&appConfig).SetRouter(app, appLogger)
	appLogger.Info("Setup application routers")

	appLogger.Info("Starting up application server on port %v", appConfig.ServerConfig.Port)
	serverErr := fasthttp.ListenAndServe(fmt.Sprintf("%s:%s", appConfig.ServerConfig.Host, appConfig.ServerConfig.Port), router.Handler)
	appLogger.Info("error while serving: %v", serverErr)
}
