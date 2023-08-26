package ghintegration

import (
	"fmt"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	logger "github.com/rs/zerolog/log"
)

var pluginTriggeredRecently = make(chan bool, 1)

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{config.FetchConfig().SERVERORIGIN + "/trigger"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"POST"})
	// ttl := handlers.MaxAge(3600)

	logger.Info().Msg(fmt.Sprintf("Starting server on port: %v", config.FetchConfig().Ports.GitHub))
	logger.Info().Err(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().Ports.GitHub), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	routeHandler := middleware.RouteHandler{}
	r := mux.NewRouter()
	r.Use(routeHandler.InternalOriginMiddleware)
	r.HandleFunc("/trigger", reqHandler).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	if <-pluginTriggeredRecently {
		logger.Info().Msg("Plugin already triggered recently")
		pluginTriggeredRecently <- true
		http.Error(w, "Plugin already triggered recently, try again after an hour", http.StatusTooManyRequests)
		w.Write([]byte("Plugin already triggered recently, try again after an hour"))
		return
	}
	go func() {
		time.Sleep(time.Hour)
		pluginTriggeredRecently <- false
	}()
	go main()
}

func StartServer() {
	pluginTriggeredRecently <- false
	handleRequests()
}
