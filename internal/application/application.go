package application

import (
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/accountmanagement"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/data"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/graphql"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/plugin"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/schedule"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Application struct {
	log                  logger.Logger
	configs              map[string]models.DBConfig
	DataSvc              ports.DataOps
	AccountManagementSvc ports.AccountOps
	PluginSvc            ports.PluginOps
	ScheduleSvc          ports.ScheduleOps
	GraphQLSvc           ports.GraphQLOps
}

func NewApplication(log logger.Logger, configs map[string]models.DBConfig) *Application {
	log.Infof("Setting up application services")

	application := &Application{
		log:     log,
		configs: configs,
	}

	application.dataService().accountManagementService().pluginService().scheduleService().build()

	return application
}

func (a *Application) dataService() *Application {
	dataSvcMongoDBConfig := a.configs[constants.DB_MONGODB]
	dataSvcMongoDBConfig.Database = []interface{}{"Graph", "GitHub"}
	dataSvcMongoDBConfig.Collection = []interface{}{} // TODO
	a.DataSvc = data.NewApplication(a.log.WithSubmodule("data"), repository.NewRepository(a.log.WithSubmodule("repository")).WithMongoDB(a.configs[constants.DB_MONGODB]).Build())
	return a
}

func (a *Application) accountManagementService() *Application {
	accountSvcMariaDBConfig := a.configs[constants.DB_MARIADB]
	accountSvcMariaDBConfig.Database = []interface{}{"Account"}
	accountSvcMariaDBConfig.Table = []interface{}{"User", "Secrets"}

	a.AccountManagementSvc = accountmanagement.NewApplication(a.log.WithSubmodule("account"), repository.NewRepository(a.log.WithSubmodule("repository")).WithMariaDB(a.configs[constants.DB_MARIADB]).Build())
	return a
}

func (a *Application) pluginService() *Application {
	pluginSvcMongoDBConfig := a.configs[constants.DB_MONGODB]
	pluginSvcMongoDBConfig.Database = []interface{}{"Chess", "GitHub"}
	pluginSvcMongoDBConfig.Collection = []interface{}{} // TODO
	pluginSvcCacheConfig := a.configs[constants.DB_REDIS]
	pluginSvcCacheConfig.Database = []interface{}{} // TODO

	a.PluginSvc = plugin.NewApplication(a.log.WithSubmodule("plugin"), repository.NewRepository(a.log.WithSubmodule("repository")).WithRedisCache(a.configs[constants.DB_REDIS]).WithMongoDB(a.configs[constants.DB_MONGODB]).Build())
	return a
}

func (a *Application) scheduleService() *Application {
	scheduleSvcMongoDBConfig := a.configs[constants.DB_MONGODB]
	scheduleSvcMongoDBConfig.Database = []interface{}{"Schedule"}
	scheduleSvcMongoDBConfig.Collection = []interface{}{} // TODO

	a.ScheduleSvc = schedule.NewApplication(a.log.WithSubmodule("schedule"), repository.NewRepository(a.log.WithSubmodule("repository")).WithMongoDB(a.configs[constants.DB_MONGODB]).Build())
	return a
}

func (a *Application) graphQLService() *Application {
	a.GraphQLSvc = graphql.NewApplication(a.log.WithSubmodule("graphql"), repository.NewRepository(a.log.WithSubmodule("repository")).Build())
	return a
}

func (a *Application) build() *Application {
	return a
}

func (a *Application) DataService() ports.DataOps {
	return a.DataSvc
}

func (a *Application) AccountManagementService() ports.AccountOps {
	return a.AccountManagementSvc
}

func (a *Application) PluginService() ports.PluginOps {
	return a.PluginSvc
}

func (a *Application) ScheduleService() ports.ScheduleOps {
	return a.ScheduleSvc
}

func (a *Application) GraphQLService() ports.GraphQLOps {
	return a.GraphQLSvc
}
