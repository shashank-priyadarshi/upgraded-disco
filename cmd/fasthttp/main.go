package main

import (
	"fmt"
	driver "github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/fasthttp"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application"

	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/config"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

func main() {
	// TODO: server will implement builder pattern to create instance with required dependencies
	appLogger := logger.NewLogger("INFO", "")
	configSource := config.Source{
		Source: "yaml",
	}
	appConfig, err := configSource.Load()
	if err != nil {
		panic(fmt.Sprintf("error initialising config from %s: %v", configSource.Source, err))
	}

	app := application.NewApplication(appLogger, appConfig.DBConfig)
	router := driver.NewRouter(&appConfig).SetRouter(app, appLogger)
	server := fasthttp.ListenAndServe("8080", router.Handler)
	appLogger.Infof("error while serving: %v", server)
}
