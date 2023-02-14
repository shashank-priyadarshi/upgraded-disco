package server

import (
	"fmt"
	"log"
	"net/http"
	"server/common"
	"server/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/biodata", returnBiodata).Methods("GET")
	r.HandleFunc("/githubdata", returnGitHubData).Methods("GET")
	r.HandleFunc("/todos", todos).Methods("POST")
	r.HandleFunc("/trigger", triggerPlugin).Methods("POST")
	r.HandleFunc("/graphql", graphqlHandler).Methods("POST")
	r.HandleFunc("/schedule", writeNewSchedule).Methods("POST")
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
