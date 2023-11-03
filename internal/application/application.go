package application

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/accountmanagement"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/data"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/graphql"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/plugin"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/schedule"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports/application/services"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports/frameworks/driven/database"
)

type Application struct {
	db          database.ServiceRepository
	dataSvc     services.DataOps
	accountSvc  services.AccountOps
	pluginSvc   services.PluginOps
	scheduleSvc services.ScheduleOps
	graphQLSvc  services.GraphQLOps
}

func NewApplication(database database.ServiceRepository) *Application {
	return &Application{
		db:          database,
		dataSvc:     data.NewApplication(database),
		accountSvc:  accountmanagement.NewApplication(database),
		pluginSvc:   plugin.NewApplication(database),
		scheduleSvc: schedule.NewApplication(database),
		graphQLSvc:  graphql.NewApplication(database),
	}
}
