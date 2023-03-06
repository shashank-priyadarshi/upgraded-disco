package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"server/auth"
	"server/common"
	"server/config"
	mongoconnection "server/db/mongo"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("endpoint %v with method %v\n", r.URL.Path, r.Method)
	// email/phone, password
	// loggedin: bool, jwt token
	rawResp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := auth.User{}
	err = user.VerifyCredentials(rawResp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("endpoint %v with method %v\n", r.URL.Path, r.Method)
	// name, email, password, phone
	// loggedin: bool, jwt token
	rawResp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := auth.User{}
	err = user.AddNewUser(rawResp)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func forgotPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("endpoint %v with method %v\n", r.URL.Path, r.Method)
	// email/phone
	// resp true/false
	// await otp
	rawResp, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := auth.User{}
	err = user.ResetPassword(rawResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	// err = auth.SendOTP(rawResp)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
}

func returnBiodata(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.BIODATA)
	w.Write(response)
}

func returnGitHubData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Endpoint Hit: %v with %v method\n", r.URL.Path, r.Method)
	response := mongoconnection.ReadDataFromCollection(config.FetchConfig().Collections.GITHUBDATA)
	w.Write(response)
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
		w.Write(response)
	}
}
