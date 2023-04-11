package server

import (
	"fmt"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/newrelic"
	logger "github.com/rs/zerolog/log"
)

func routes(app *newrelic.Application) *mux.Router {
	routeHandler := middleware.RouteHandler{}
	r := mux.NewRouter()
	// r.Use(middleware.ExternalOriginMiddleware)
	// r.Use(middleware.AddResponseHeaders)
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/biodata", returnBiodata)).Methods("GET")
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/graphdata", returnGraphData)).Methods("GET")

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/graphql", graphqlHandler)).Methods("POST")
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/credentials", credentials)).Methods("POST")
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/schedule", writeNewSchedule)).Methods("POST")

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/githubdata", returnGitHubData)).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(returnGitHubData)))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/todos", todos)).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(todos)))
	r.HandleFunc(newrelic.WrapHandleFunc(app, "/trigger", triggerPlugin)).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(triggerPlugin)))

	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func handleRequests(app *newrelic.Application) {
	router := routes(app)

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	// ttl := handlers.MaxAge(3600)

	logger.Info().Msg(fmt.Sprintf("Starting server on port: %v", config.FetchConfig().Ports.Server))
	logger.Info().Err(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().Ports.Server), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func StartServer() {
	app, err := newrelicObs()
	if err != nil {
		logger.Info().Err(err).Msg("error while setting up application observability using newrelic: ")
	}
	handleRequests(app)
}

func newrelicObs() (app *newrelic.Application, err error) {
	app, err = newrelic.NewApplication(
		newrelic.ConfigAppName(config.FetchConfig().Application),
		newrelic.ConfigLicense(config.FetchConfig().License),
		newrelic.ConfigAppLogForwardingEnabled(config.FetchConfig().LogForward),
	)
	return
}
