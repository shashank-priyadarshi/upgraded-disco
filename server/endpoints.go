package server

import (
	"fmt"
	"log"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	routeHandler := middleware.RouteHandler{}
	r := mux.NewRouter()
	// r.Use(middleware.ExternalOriginMiddleware)
	// r.Use(middleware.AddResponseHeaders)
	r.HandleFunc("/biodata", returnBiodata).Methods("GET")
	r.HandleFunc("/graphdata", returnGraphData).Methods("GET")

	r.HandleFunc("/graphql", graphqlHandler).Methods("POST")
	r.HandleFunc("/credentials", credentials).Methods("POST")

	r.HandleFunc("/githubdata", returnGitHubData).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(returnGitHubData)))
	r.HandleFunc("/todos", todos).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(todos)))
	r.HandleFunc("/trigger", triggerPlugin).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(triggerPlugin)))
	r.HandleFunc("/schedule", writeNewSchedule).Methods("POST").Handler(routeHandler.AuthMiddleware(http.HandlerFunc(writeNewSchedule)))

	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	// ttl := handlers.MaxAge(3600)

	fmt.Println("Starting server on port: ", config.FetchConfig().SERVERPORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().SERVERPORT), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func StartServer() {
	handleRequests()
}
