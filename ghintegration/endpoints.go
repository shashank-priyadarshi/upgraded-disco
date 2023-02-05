package ghintegration

import (
	"fmt"
	"log"
	"net/http"
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

	fmt.Println("Starting server on port: ", config.FetchConfig().GHINTEGRATIONORIGIN)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().GHINTEGRATIONORIGIN), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	r := mux.NewRouter()
	// r.Use(middleware.OriginMiddleware)
	r.HandleFunc("/trigger", reqHandler).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(invalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	go main()
}

func invalidEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	http.Error(w, "Endpoint does not exist", http.StatusNotFound)
}

func StartServer() {
	handleRequests()
}
