package ghintegration

import (
	"fmt"
	"log"
	"net/http"
	"server/common"
	"server/config"
	"server/middleware"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var pluginTriggeredRecently = make(chan bool, 1)

func handleRequests() {
	router := routes()

	credentials := handlers.AllowCredentials()
	origins := handlers.AllowedOrigins([]string{config.FetchConfig().SERVERORIGIN + "/trigger"})
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Referrer-Policy"})
	methods := handlers.AllowedMethods([]string{"POST", "OPTIONS"})
	//ttl := handlers.MaxAge(3600)

	fmt.Println("Starting server on port: ", config.FetchConfig().GHINTEGRATIONPORT)
	log.Println(http.ListenAndServe(fmt.Sprintf(":%v", config.FetchConfig().GHINTEGRATIONPORT), handlers.CORS(credentials, headers, methods, origins)(router)))
}

func routes() *mux.Router {
	r := mux.NewRouter()
	r.Use(middleware.OriginMiddleware)
	r.HandleFunc("/trigger", reqHandler).Methods("POST")
	r.NotFoundHandler = http.HandlerFunc(common.InvalidEndpoint)
	return r
}

func reqHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	if <-pluginTriggeredRecently {
		fmt.Println("Plugin already triggered recently")
		http.Error(w, "Plugin already triggered recently, try again after an hour", http.StatusTooManyRequests)
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
