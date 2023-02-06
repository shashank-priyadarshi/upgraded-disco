package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/common"
	"server/config"
	"server/mongoconnection"
)

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.BIODATA)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("error while encoding request data for endpoint %v: %v\n", r.URL.Path, err)
	}
}

func returnGitHubData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.GITHUBDATA)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Printf("error while encoding request data for endpoint %v: %v\n", r.URL.Path, err)
	}
}

func writeNewSchedule(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	// response := mongoconnection.WriteDataToCollection(config.FetchConfig().Collections.SCHEDULE, r)
	// json.NewEncoder(w).Encode(response)
}

func triggerPlugin(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response, statusCode := common.NoAuthAPICall(fmt.Sprintf("http://localhost:%v/trigger", config.FetchConfig().GHINTEGRATIONPORT), fmt.Sprintf("%v/trigger", config.FetchConfig().SERVERORIGIN), []byte(""))
	if statusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error while triggering plugin: %v", string(response)), statusCode)
	} else {
		_, err := w.Write([]byte("Plugin triggered successfully!"))
		if err != nil {
			fmt.Printf("error while writing response for endpoint %v: %v\n", r.URL.Path, err)
		}
	}
}

func todos(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error while reading request body for endpoint %v: %v\n", r.URL.Path, err)
	}
	response, statusCode := common.NoAuthAPICall(fmt.Sprintf("http://localhost:%v/list", config.FetchConfig().TODOAPIPORT), fmt.Sprintf("%v/todos", config.FetchConfig().SERVERORIGIN), body)
	if statusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error while fetching todos list: %v", string(response)), statusCode)
	} else {
		err := json.NewEncoder(w).Encode(string(response))
		if err != nil {
			fmt.Printf("error while encoding request data for endpoint %v: %v\n", r.URL.Path, err)
		}
	}
}
