package application

import (
	"github.com/shashank-priyadarshi/upgraded-disco/constants"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/adapters/repository"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/accountmanagement"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/data"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/graphql"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/plugin"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/schedule"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

func NewApplication(log logger.Logger, configs map[string]models.DBConfig) *models.Application {
	log.Infof("Setting up application services")

	dataSvcMongoDBConfig := configs[constants.DB_MONGODB]
	dataSvcMongoDBConfig.Database = []interface{}{"Graph", "GitHub"}
	dataSvcMongoDBConfig.Collection = []interface{}{} // TODO

	accountSvcMariaDBConfig := configs[constants.DB_MARIADB]
	accountSvcMariaDBConfig.Database = []interface{}{"Account"}
	accountSvcMariaDBConfig.Table = []interface{}{"User", "Secrets"}

	pluginSvcMongoDBConfig := configs[constants.DB_MONGODB]
	pluginSvcMongoDBConfig.Database = []interface{}{"Chess", "GitHub"}
	pluginSvcMongoDBConfig.Collection = []interface{}{} // TODO
	pluginSvcCacheConfig := configs[constants.DB_REDIS]
	pluginSvcCacheConfig.Database = []interface{}{} // TODO

	scheduleSvcMongoDBConfig := configs[constants.DB_MONGODB]
	scheduleSvcMongoDBConfig.Database = []interface{}{"Schedule"}
	scheduleSvcMongoDBConfig.Collection = []interface{}{} // TODO

	// All services will have their own table & collection names
	return &models.Application{
		DataSvc:     data.NewApplication(log.WithSubmodule("data"), repository.NewRepository(log.WithSubmodule("repository")).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		AccountSvc:  accountmanagement.NewApplication(log.WithSubmodule("account"), repository.NewRepository(log.WithSubmodule("repository")).WithMariaDB(configs[constants.DB_MARIADB]).Build()),
		PluginSvc:   plugin.NewApplication(log.WithSubmodule("plugin"), repository.NewRepository(log.WithSubmodule("repository")).WithRedisCache(configs[constants.DB_REDIS]).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		ScheduleSvc: schedule.NewApplication(log.WithSubmodule("schedule"), repository.NewRepository(log.WithSubmodule("repository")).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		GraphQLSvc:  graphql.NewApplication(log.WithSubmodule("graphql"), repository.NewRepository(log.WithSubmodule("repository")).Build()),
	}
}
