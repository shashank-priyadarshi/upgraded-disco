package application

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/accountmanagement"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/data"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/graphql"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/plugin"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/application/services/schedule"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
)

type Application struct {
	db          ports.ServiceRepository
	dataSvc     ports.DataOps
	accountSvc  ports.AccountOps
	pluginSvc   ports.PluginOps
	scheduleSvc ports.ScheduleOps
	graphQLSvc  ports.GraphQLOps
}

func NewApplication(database ports.ServiceRepository) *Application {
	return &Application{
		db:          database,
		dataSvc:     data.NewApplication(database),
		accountSvc:  accountmanagement.NewApplication(database),
		pluginSvc:   plugin.NewApplication(database),
		scheduleSvc: schedule.NewApplication(database),
		graphQLSvc:  graphql.NewApplication(database),
	}
}
