package form

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	logger "github.com/rs/zerolog/log"
)

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{config.FetchConfig().SERVERORIGIN})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"POST"})
	// ttl := handlers.MaxAge(3600)

	logger.Info().Msg(fmt.Sprintf("Starting server on port: %v", config.FetchConfig().Ports.Schedule))
	logger.Info().Err(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().Ports.Schedule), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	routeHandler := middleware.RouteHandler{}
	r := mux.NewRouter()
	r.Use(routeHandler.InternalOriginMiddleware)
	r.HandleFunc("/new", reqHandler).Methods("POST") // add message to mongo, forward email to email address
	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	var body *Message
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := body.forwardMessage()
	if err != nil {
		logger.Info().Err(err).Msg("error while forwarding message: ")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	if body.Newsletter {
		err = body.addToNewsletter()
		if err != nil {
			logger.Info().Err(err).Msg("error while adding to newsletter: ")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	if body.Meeting {
		err = body.scheduleMeet()
		if err != nil {
			logger.Info().Err(err).Msg("error while scheduling meeting: ")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
}

func StartServer() {
	handleRequests()
}
