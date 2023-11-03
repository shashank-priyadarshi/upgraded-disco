package fasthttp

import (
	"fmt"

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

func (r *Router) SetRouter(services interface{}, log logger.Logger) *router.Router {
	r.setRouterConfig()
	r.setDataRouter(services)
	r.setAccountRouter(services)
	r.setPluginRouter(services)
	r.setScheduleRouter(services)
	r.setGraphQLRouter(services)
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

func (r *Router) setDataRouter(services interface{}) *router.Router {
	// github: auth
	dataGroup := r.setV1Router("/data")
	{
		dataGroup.POST("/graph", nil)
		dataGroup.POST("/github", nil)
	}
	return r.router
}

func (r *Router) setAccountRouter(services interface{}) *router.Router {
	accountGroup := r.setV1Router("/account")
	{
		accountGroup.POST("/signup", nil)
		accountGroup.POST("/login", nil)
		accountGroup.POST("/reset-password", nil)
		accountGroup.POST("/unregister", nil)
	}
	return r.router
}

func (r *Router) setPluginRouter(services interface{}) *router.Router {
	// list: auth
	// update: auth
	// install: auth
	// trigger: auth
	pluginGroup := r.setV1Router("/plugin")
	{
		pluginGroup.POST("/list", nil)
		pluginGroup.POST("/update", nil)
		pluginGroup.POST("/install", nil)
		pluginGroup.POST("/trigger", nil)
	}
	return r.router
}

func (r *Router) setScheduleRouter(services interface{}) *router.Router {
	// /list: auth
	scheduleGroup := r.setV1Router("/schedule")
	{
		scheduleGroup.POST("/list", nil)
		scheduleGroup.POST("/create", nil)
		scheduleGroup.POST("/cancel", nil)
	}
	return r.router
}

func (r *Router) setGraphQLRouter(services interface{}) *router.Router {
	// TODO
	// {
	// 	v1.POST("/graphql", nil)
	// }
	return r.router
}
