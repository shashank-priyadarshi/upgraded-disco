package config

import (
	"encoding/base64"
	"fmt"
	"os"
)

type Configuration struct {
	SQLURI            string
	MongoURI          string
	DBNAME            string
	SERVERPORT        string
	TODOAPIPORT       string
	SERVERORIGIN      string
	GHINTEGRATIONPORT string
	GITHUBTOKEN       string
	GITHUBUSERNAME    string
	ALLOWEDORIGIN     string
	SECRETKEY         []byte
	Collections
}

type Collections struct {
	BIODATA    string
	GITHUBDATA string
	TODOS      string
	GRAPHDATA  string
	SCHEDULE   string
}

func FetchConfig() Configuration {
	return Configuration{
		DBNAME:            os.Getenv("DB_NAME"),
		SQLURI:            os.Getenv("SQL_URI"),
		MongoURI:          os.Getenv("MONGO_URI"),
		SERVERPORT:        os.Getenv("SERVER_PORT"),
		GITHUBTOKEN:       os.Getenv("GITHUB_TOKEN"),
		TODOAPIPORT:       os.Getenv("TODO_API_PORT"),
		ALLOWEDORIGIN:     os.Getenv("ALLOWED_ORIGIN"),
		GITHUBUSERNAME:    os.Getenv("GITHUB_USERNAME"),
		GHINTEGRATIONPORT: os.Getenv("GH_INTEGRATION_ORIGIN"),
		SERVERORIGIN:      fmt.Sprintf("http://localhost:%v", os.Getenv("SERVER_PORT")),
		SECRETKEY:         fetchSecretKey(),
		Collections: Collections{
			BIODATA:    os.Getenv("BIO"),
			GITHUBDATA: os.Getenv("GITHUB"),
			TODOS:      os.Getenv("TODOS"),
			GRAPHDATA:  os.Getenv("GRAPH"),
			SCHEDULE:   os.Getenv("SCHEDULE"),
		},
	}
}

func fetchSecretKey() (key []byte) {
	// Read the secret key from the environment variable
	encodedKey := os.Getenv("SECRET_KEY")
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		fmt.Printf("error while decoding secret key: %v", err)
		return []byte("")
	}
	return
}
