package config

import (
	"fmt"
	"os"
)

type Configuration struct {
	MongoURI            string
	DBNAME              string
	SERVERPORT          string
	TODOAPIPORT         string
	SERVERORIGIN        string
	GHINTEGRATIONORIGIN string
	GITHUBTOKEN         string
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
		MongoURI:            os.Getenv("MONGO_URI"),
		DBNAME:              os.Getenv("DB_NAME"),
		SERVERPORT:          os.Getenv("SERVER_PORT"),
		TODOAPIPORT:         os.Getenv("TODO_API_PORT"),
		GITHUBTOKEN:         os.Getenv("GITHUB_TOKEN"),
		GHINTEGRATIONORIGIN: os.Getenv("GH_INTEGRATION_ORIGIN"),
		SERVERORIGIN:        fmt.Sprintf("http://localhost:%v", os.Getenv("SERVER_PORT")),
		Collections: Collections{
			BIODATA:    os.Getenv("COLLECTION_BIODATA"),
			GITHUBDATA: os.Getenv("COLLECTION_GITHUBDATA"),
			TODOS:      os.Getenv("COLLECTION_TODOS"),
			ARTICLES:   os.Getenv("COLLECTION_ARTICLES"),
			SCHEDULE:   os.Getenv("COLLECTION_SCHEDULE"),
		},
	}
}
