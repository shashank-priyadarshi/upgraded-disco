package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/auth"
	"server/common"
	"server/config"
	mongoconnection "server/db/mongo"

	logger "github.com/rs/zerolog/log"
)

func credentials(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("endpoint %v with method %v\n", r.URL.Path, r.Method))
	// name, email, password, phone
	// loggedin: bool, jwt token
	rawResp, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Info().Err(err).Msg("error while reading request body: ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := auth.User{}
	token, err := user.ParseCredentials(rawResp)
	if err != nil {
		logger.Info().Err(err).Msg("error while parsing credentials: ")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(struct {
		Token string `json:"token"`
	}{token})
	if err != nil {
		logger.Info().Err(err).Msg("failed to marshal token")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.BIODATA)
	w.Write(response)
}

func returnGitHubData(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.GITHUBDATA)
	w.Write(response)
}

func returnGraphData(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.GRAPHDATA)
	w.Write(response)
}

func writeNewSchedule(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response, statusCode := common.NoAuthAPICall(fmt.Sprintf("http://localhost:%v/trigger", config.FetchConfig().Ports.GitHub), fmt.Sprintf("%v/trigger", config.FetchConfig().SERVERORIGIN), []byte(""))
	if statusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error while triggering plugin: %v", string(response)), statusCode)
	} else {
		_, err := w.Write([]byte("Plugin triggered successfully!"))
		if err != nil {
			logger.Info().Msg(fmt.Sprintf("error while writing response for endpoint %v: %v\n", r.URL.Path, err))
		}
	}
}

func triggerPlugin(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	response, statusCode := common.NoAuthAPICall(fmt.Sprintf("http://localhost:%v/trigger", config.FetchConfig().Ports.GitHub), fmt.Sprintf("%v/trigger", config.FetchConfig().SERVERORIGIN), []byte(""))
	if statusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error while triggering plugin: %v", string(response)), statusCode)
	} else {
		_, err := w.Write([]byte("Plugin triggered successfully!"))
		if err != nil {
			logger.Info().Msg(fmt.Sprintf("error while writing response for endpoint %v: %v\n", r.URL.Path, err))
		}
	}
}

func todos(w http.ResponseWriter, r *http.Request) {
	logger.Info().Msg(fmt.Sprintf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method))
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Info().Msg(fmt.Sprintf("error while reading request body from endpoint %v: %v\n", r.URL.Path, err))
	}
	response, statusCode := common.NoAuthAPICall(fmt.Sprintf("http://localhost:%v/list", config.FetchConfig().Ports.Todos), fmt.Sprintf("%v/todos", config.FetchConfig().SERVERORIGIN), body)
	if statusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Error while fetching todos list: %v", string(response)), statusCode)
	} else {
		w.Write(response)
	}
}
