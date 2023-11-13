package fasthttp

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"

	"github.com/fasthttp/router"
	"github.com/shashank-priyadarshi/upgraded-disco/models"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
)

type Router struct {
	router *router.Router
	config *models.Config
}

func NewRouter(config *models.Config) *Router {
	return &Router{
		router: router.New(),
		config: config,
	}
}

func (r *Router) SetRouter(services *models.Application, log logger.Logger) *router.Router {
	r.setRouterConfig()
	r.setDataRouter(services.DataSvc, log)
	r.setAccountRouter(services.AccountSvc, log)
	r.setPluginRouter(services.PluginSvc, log)
	r.setScheduleRouter(services.ScheduleSvc, log)
	r.setGraphQLRouter(services.GraphQLSvc, log)
	r.setRouterConfig()
	return r.router
}

func (r *Router) setRouterConfig() *router.Router {
	// TODO: Setup all configs:
	// Origin: dev, ssnk.in
	// Headers: "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"
	// Methods: "POST", "OPTIONS"
	return r.router
}

func (r *Router) setV1Router(group string) *router.Group {
	return r.router.Group("/v1").Group(fmt.Sprintf("/%s", group))
}

func (r *Router) setDataRouter(ops ports.DataOps, log logger.Logger) *router.Router {
	// github: auth
	handler := Data{DataOps: ops, Logger: log}
	dataGroup := r.setV1Router("/data")
	{
		dataGroup.POST("/graph", handler.GetGraphData)
		dataGroup.POST("/github", handler.GetGitHubData)
	}
	return r.router
}

func (r *Router) setAccountRouter(ops ports.AccountOps, log logger.Logger) *router.Router {
	handler := Account{AccountOps: ops, Logger: log}
	accountGroup := r.setV1Router("/account")
	{
		accountGroup.POST("/signup", handler.RegisterUser)
		accountGroup.POST("/login", handler.Login)
		accountGroup.POST("/reset-password", handler.ResetPassword)
		accountGroup.POST("/unregister", handler.DeleteUser)
	}
	return r.router
}

func (r *Router) setPluginRouter(ops ports.PluginOps, log logger.Logger) *router.Router {
	// list: auth
	// update: auth
	// install: auth
	// trigger: auth
	handler := Plugins{PluginOps: ops, Logger: log}
	pluginGroup := r.setV1Router("/plugins")
	{
		pluginGroup.POST("/list", handler.List)
		pluginGroup.POST("/update", handler.Update)
		pluginGroup.POST("/install", handler.Install)
		pluginGroup.POST("/trigger", handler.Trigger)
	}
	return r.router
}

func (r *Router) setScheduleRouter(ops ports.ScheduleOps, log logger.Logger) *router.Router {
	// /list: auth
	handler := Schedule{ScheduleOps: ops, Logger: log}
	scheduleGroup := r.setV1Router("/schedule")
	{
		scheduleGroup.POST("/list", handler.List)
		scheduleGroup.POST("/create", handler.Create)
		scheduleGroup.POST("/cancel", handler.Delete)
	}
	return r.router
}

func (r *Router) setGraphQLRouter(ops ports.GraphQLOps, log logger.Logger) *router.Router {
	// TODO
	// handler := graphql.GraphQL{GraphQLOps: ops, Logger: log}
	// {
	// 	v1.POST("/graphql", handler.GraphQL())
	// }
	return r.router
}
