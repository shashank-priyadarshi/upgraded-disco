package server

import (
	"fmt"
	"log"
	"net/http"
	"server/config"
	"server/todos"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/graphql", graphqlHandler).Methods("POST")
	r.HandleFunc("/trigger", invalidEndpoint).Methods("POST")
	r.HandleFunc("/todos", todos.ReturnTodos).Methods("GET")
	r.HandleFunc("/biodata", returnBiodata).Methods("GET")
	r.HandleFunc("/schedule", writeNewSchedule).Methods("POST")
	r.HandleFunc("/githubdata", returnGitHubData).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(invalidEndpoint)
	return r
}

//func originValidator(origin string) bool {
//	valid := false
//	err := pool.QueryRow("SELECT IF(origin=?, True, False) as 'valid' FROM origins", origin).Scan(&valid)
//	if err != nil {
//		return false
//	}
//	return valid
//}

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{"*"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	//ttl := handlers.MaxAge(3600)

	fmt.Println("Starting server on port: ", config.FetchConfig().SERVERPORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().SERVERPORT), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func StartServer() {
	handleRequests()
}
