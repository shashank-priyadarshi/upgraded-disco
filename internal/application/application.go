package application

import (
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/databases"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/accountmanagement"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/data"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/graphql"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/plugin"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/schedule"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/pubsub"
)

type Application struct {
	log                  logger.Logger
	dbConfigs            map[string]models.DBConfig
	pluginConfig         *models.PluginsConfig
	DataSvc              ports.DataOps
	AccountManagementSvc ports.AccountOps
	PluginSvc            ports.PluginOps
	ScheduleSvc          ports.ScheduleOps
	GraphQLSvc           ports.GraphQLOps
	BrokerSvc            pubsub.PubSub
}

func NewApplication(log logger.Logger, configs map[string]models.DBConfig, pluginConfig *models.PluginsConfig) *Application {
	log.Infof("Setting up application services")

	application := &Application{
		log:          log,
		dbConfigs:    configs,
		pluginConfig: pluginConfig,
	}

	application.brokerService().dataService().accountManagementService().pluginService().scheduleService().build()

	return application
}

func (a *Application) dataService() *Application {
	dataSvcMongoDBConfig := a.dbConfigs[constants.DB_MONGODB]
	dataSvcMongoDBConfig.Database = []interface{}{"Graph", "GitHub"}
	dataSvcMongoDBConfig.Collection = []interface{}{} // TODO
	a.DataSvc = data.NewApplication(a.log.WithSubmodule("data"), databases.NewRepository(a.log.WithSubmodule("repository")).WithMongoDB(a.dbConfigs[constants.DB_MONGODB]).Build())
	return a
}

func (a *Application) accountManagementService() *Application {
	accountSvcMariaDBConfig := a.dbConfigs[constants.DB_MARIADB]
	accountSvcMariaDBConfig.Database = []interface{}{"Account"}
	accountSvcMariaDBConfig.Table = []interface{}{"User", "Secrets"}

	a.AccountManagementSvc = accountmanagement.NewApplication(a.log.WithSubmodule("account"), databases.NewRepository(a.log.WithSubmodule("repository")).WithMariaDB(a.dbConfigs[constants.DB_MARIADB]).Build())
	return a
}

func (a *Application) pluginService() *Application {
	pluginSvcMongoDBConfig := a.dbConfigs[constants.DB_MONGODB]
	pluginSvcMongoDBConfig.Database = []interface{}{"Chess", "GitHub"}
	pluginSvcMongoDBConfig.Collection = []interface{}{} // TODO
	pluginSvcCacheConfig := a.dbConfigs[constants.DB_REDIS]
	pluginSvcCacheConfig.Database = []interface{}{} // TODO

	a.PluginSvc = plugin.NewApplication(a.log.WithSubmodule("plugin"), databases.NewRepository(a.log.WithSubmodule("repository")).WithRedisCache(a.dbConfigs[constants.DB_REDIS]).WithMongoDB(a.dbConfigs[constants.DB_MONGODB]).Build(), a.BrokerSvc, a.pluginConfig)
	return a
}

func (a *Application) scheduleService() *Application {
	scheduleSvcMongoDBConfig := a.dbConfigs[constants.DB_MONGODB]
	scheduleSvcMongoDBConfig.Database = []interface{}{"Schedule"}
	scheduleSvcMongoDBConfig.Collection = []interface{}{} // TODO

	a.ScheduleSvc = schedule.NewApplication(a.log.WithSubmodule("schedule"), databases.NewRepository(a.log.WithSubmodule("repository")).WithMongoDB(a.dbConfigs[constants.DB_MONGODB]).Build())
	return a
}

func (a *Application) graphQLService() *Application {
	a.GraphQLSvc = graphql.NewApplication(a.log.WithSubmodule("graphql"), databases.NewRepository(a.log.WithSubmodule("repository")).Build())
	return a
}

func (a *Application) brokerService() *Application {
	if broker, err := pubsub.NewBroker(); err != nil {
		a.log.Errorf("Error initializing broker service: ", err)
	} else {
		a.BrokerSvc = broker
	}
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
