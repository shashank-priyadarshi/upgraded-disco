package fasthttp

import (
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

// TODO: setup auth middleware
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

func (r *Router) setV1Router() *router.Group {
	return r.router.Group("/v1")
}

func (r *Router) setDataRouter(ops ports.DataOps, log logger.Logger) *router.Router {
	// github: auth
	handler := Data{DataOps: ops, Logger: log}
	dataGroup := r.setV1Router().Group("/data")
	{
		dataGroup.GET("/graph", handler.Graph)
		dataGroup.GET("/github", handler.GitHub)
	}
	return r.router
}

func (r *Router) setAccountRouter(ops ports.AccountOps, log logger.Logger) *router.Router {
	handler := Account{AccountOps: ops, Logger: log}
	accountGroup := r.setV1Router().Group("/account")
	{
		accountGroup.POST("/signup", handler.Register)
		accountGroup.POST("/login", handler.Login)
		accountGroup.PUT("/:id", handler.Update)
		accountGroup.DELETE("/:id", handler.Delete)
	}
	return r.router
}

func (r *Router) setPluginRouter(ops ports.PluginOps, log logger.Logger) *router.Router {
	// list: auth
	// update: auth
	// install: auth
	// trigger: auth
	handler := Plugins{PluginOps: ops, Logger: log}
	pluginGroup := r.setV1Router().Group("/plugins")
	{
		pluginGroup.POST("/", handler.Install)
		pluginGroup.GET("/", handler.Get)
		pluginGroup.GET("/:id", handler.Get)
		pluginGroup.PUT("/:id", handler.Update)
		pluginGroup.POST("/:id/trigger", handler.Trigger)
		pluginGroup.DELETE("/:id", handler.Delete)
	}
	return r.router
}

func (r *Router) setScheduleRouter(ops ports.ScheduleOps, log logger.Logger) *router.Router {
	// list: auth
	handler := Schedule{ScheduleOps: ops, Logger: log}
	scheduleGroup := r.setV1Router().Group("/schedule")
	{
		scheduleGroup.POST("/", handler.Create)
		scheduleGroup.GET("/", handler.Get)
		scheduleGroup.GET("/:id", handler.Get)
		scheduleGroup.PUT("/:id", handler.Update)
		scheduleGroup.DELETE("/:id", handler.Delete)
	}
	return r.router
}

func (r *Router) setGraphQLRouter(ops ports.GraphQLOps, log logger.Logger) *router.Router {
	//handler := GraphQL{GraphQLOps: ops, Logger: log}
	//v1 := r.setV1Router("/graphql")
	//{
	//v1.POST("*", handler.GraphQL)
	//}
	return r.router
}
