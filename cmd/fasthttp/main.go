package main

import (
	"fmt"

	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/config"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driven/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
	driver "github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/frameworks/driver/fasthttp"
)

func main() {
	// TODO: server will implement builder pattern to create instance with required dependencies
	logger := logger.NewLogger("INFO", "")
	configSource := config.Source{
		Source: "yaml",
	}
	config, err := configSource.Load()
	if err != nil {
		panic(fmt.Sprintf("error initialising config from %s: %v", configSource.Source, err))
	}
	cache, err := databases.NewDatabase(constants.DB_REDIS, logger, config)
	if err != nil {
		panic(fmt.Sprintf("error initialising cache: %v", err))
	}
	mariadb, err := databases.NewDatabase(constants.DB_MARIADB, logger, config)
	if err != nil {
		panic(fmt.Sprintf("error initialising maridb: %v", err))
	}
	mongodb, err := databases.NewDatabase(constants.DB_MONGODB, logger, config)
	if err != nil {
		panic(fmt.Sprintf("error initialising mongodb: %v", err))
	}
	// TODO: initialise db adapters to create db repositories for all actions
	// TODO: initialise application with respective db adapters
	// setup router
	router := driver.NewRouter(&config).SetRouter(nil, logger)
	server := fasthttp.ListenAndServe("8080", router.Handler)
	logger.Infof("error while serving: %v", server)
}
