package fasthttp

import (
	"fmt"
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"

	"github.com/fasthttp/router"
	models "github.com/shashank-priyadarshi/upgraded-disco/models"
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
	r.setDataRouter(services.DataSvc)
	r.setAccountRouter(services.AccountSvc)
	r.setPluginRouter(services.PluginSvc)
	r.setScheduleRouter(services.ScheduleSvc)
	r.setGraphQLRouter(services.GraphQLSvc)
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

func (r *Router) setDataRouter(ops ports.DataOps) *router.Router {
	// github: auth
	dataGroup := r.setV1Router("/data")
	{
		dataGroup.POST("/graph", ops.GetGraphData)
		dataGroup.POST("/github", ops.GetGitHubData)
	}
	return r.router
}

func (r *Router) setAccountRouter(ops ports.AccountOps) *router.Router {
	accountGroup := r.setV1Router("/account")
	{
		accountGroup.POST("/signup", ops.RegisterUser)
		accountGroup.POST("/login", ops.Login)
		accountGroup.POST("/reset-password", ops.ResetPassword)
		accountGroup.POST("/unregister", ops.DeleteUser)
	}
	return r.router
}

func (r *Router) setPluginRouter(ops ports.PluginOps) *router.Router {
	// list: auth
	// update: auth
	// install: auth
	// trigger: auth
	pluginGroup := r.setV1Router("/plugin")
	{
		pluginGroup.POST("/list", ops.List)
		pluginGroup.POST("/update", ops.Update)
		pluginGroup.POST("/install", ops.Install)
		pluginGroup.POST("/trigger", ops.Trigger)
	}
	return r.router
}

func (r *Router) setScheduleRouter(ops ports.ScheduleOps) *router.Router {
	// /list: auth
	scheduleGroup := r.setV1Router("/schedule")
	{
		scheduleGroup.POST("/list", ops.List)
		scheduleGroup.POST("/create", ops.Create)
		scheduleGroup.POST("/cancel", ops.Delete)
	}
	return r.router
}

func (r *Router) setGraphQLRouter(ops ports.GraphQLOps) *router.Router {
	// TODO
	// {
	// 	v1.POST("/graphql", ops.GraphQL())
	// }
	return r.router
}
