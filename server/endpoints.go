package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/config"
	"server/mongoconnection"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	reqStatus := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.BIODATA)
	json.NewEncoder(w).Encode(reqStatus)
}

func returnGitHubData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	reqStatus := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.GITHUBDATA)
	json.NewEncoder(w).Encode(reqStatus)
}

func writeNewSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	reqStatus := mongoconnection.WriteDataToCollection(config.FetchConfig().Collections.SCHEDULE, r)
	json.NewEncoder(w).Encode(reqStatus)
}

func invalidEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	http.Error(w, "Endpoint does not exist", http.StatusNotFound)
}
