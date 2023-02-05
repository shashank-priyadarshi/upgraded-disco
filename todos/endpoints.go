package todos

import (
	"fmt"
	"log"
	"net/http"
	"server/common"
	"server/config"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{config.FetchConfig().SERVERORIGIN})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"POST", "OPTIONS"})
	//ttl := handlers.MaxAge(3600)

	fmt.Println("Starting server on port: ", config.FetchConfig().TODOAPIPORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().TODOAPIPORT), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	r := mux.NewRouter()
	// r.Use(middleware.OriginMiddleware)
	r.HandleFunc("/list", ReturnTodos).Methods("GET")
	r.HandleFunc("/new", reqHandler).Methods("POST")  // add todo to mongo, add new issue to respective repo
	r.HandleFunc("/done", reqHandler).Methods("POST") // delete todo from mongo, change status on github
	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
}

func StartServer() {
	handleRequests()
}
