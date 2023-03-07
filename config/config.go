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
	ARTICLES   string
	SCHEDULE   string
}

func FetchConfig() Configuration {
	return Configuration{
		SQLURI:            os.Getenv("SQL_URI"),
		DBNAME:            os.Getenv("DB_NAME"),
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
			BIODATA:    os.Getenv("COLLECTION_BIODATA"),
			GITHUBDATA: os.Getenv("COLLECTION_GITHUBDATA"),
			TODOS:      os.Getenv("COLLECTION_TODOS"),
			ARTICLES:   os.Getenv("COLLECTION_ARTICLES"),
			SCHEDULE:   os.Getenv("COLLECTION_SCHEDULE"),
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
