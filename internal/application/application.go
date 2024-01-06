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

	return &models.Application{
		DataSvc:     data.NewApplication(log, repository.NewRepository(log).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		AccountSvc:  accountmanagement.NewApplication(log, repository.NewRepository(log).WithMariaDB(configs[constants.DB_MARIADB]).Build()),
		PluginSvc:   plugin.NewApplication(log, repository.NewRepository(log).WithRedisCache(configs[constants.DB_REDIS]).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		ScheduleSvc: schedule.NewApplication(log, repository.NewRepository(log).WithMongoDB(configs[constants.DB_MONGODB]).Build()),
		GraphQLSvc:  graphql.NewApplication(log, repository.NewRepository(log).Build()),
	}
}
