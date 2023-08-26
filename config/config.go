package config

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	logger "github.com/rs/zerolog/log"
)

type Configuration struct {
	SQLURI         string
	MongoURI       string
	DBNAME         string
	SERVERORIGIN   string
	ALLOWEDORIGIN  string
	GITHUBTOKEN    string
	GITHUBUSERNAME string
	SECRETKEY      []byte
	Ports
	NewRelic
	Collections
}
type Ports struct {
	Server string
	GitHub string
}
type NewRelic struct {
	Application string
	License     string
	LogForward  bool
}
type Collections struct {
	GITHUBDATA string
	GRAPHDATA  string
}

func FetchConfig() Configuration {
	if strings.EqualFold("0", os.Getenv("SETUP")) {
		return Configuration{
			SQLURI:         "root@tcp(mysql:3306)/db",
			MongoURI:       "mongodb://mongodb:27017/test",
			DBNAME:         "test",
			SERVERORIGIN:   "http://localhost:4200",
			GITHUBTOKEN:    os.Getenv("GH"),
			GITHUBUSERNAME: "shashank-priyadarshi",
			ALLOWEDORIGIN:  "http://localhost:4200",
			SECRETKEY:      fetchSecretKey(),
			Ports: Ports{
				Server: "8085",
				GitHub: "8086",
			},
			Collections: Collections{
				GITHUBDATA: "g",
				GRAPHDATA:  "gr",
			},
			NewRelic: NewRelic{
				Application: "",
				License:     "",
				LogForward:  false,
			},
		}
	}
	return Configuration{
		DBNAME:         os.Getenv("DB_NAME"),
		SQLURI:         os.Getenv("SQL_URI"),
		MongoURI:       os.Getenv("MONGO_URI"),
		GITHUBTOKEN:    os.Getenv("GITHUB_TOKEN"),
		ALLOWEDORIGIN:  os.Getenv("ALLOWED_ORIGIN"),
		GITHUBUSERNAME: os.Getenv("GITHUB_USERNAME"),
		SERVERORIGIN:   fmt.Sprintf("http://localhost:%v", os.Getenv("SERVER_PORT")),
		SECRETKEY:      fetchSecretKey(),
		Ports: Ports{
			Server: os.Getenv("SERVER_PORT"),
			GitHub: os.Getenv("GITHUB_PORT"),
		},
		Collections: Collections{
			GITHUBDATA: os.Getenv("GITHUB"),
			GRAPHDATA:  os.Getenv("GRAPH"),
		},
		NewRelic: NewRelic{
			Application: os.Getenv("NEWRELIC_APP"),
			License:     os.Getenv("NEWRELIC_LICENSE"),
			LogForward:  os.Getenv("NEWRELIC_LOG_FORWARD") == "true",
		},
	}
}

func fetchSecretKey() (key []byte) {
	// Read the secret key from the environment variable
	encodedKey := os.Getenv("SECRET_KEY")
	key, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		logger.Info().Err(err).Msg("failed to decode secret key: ")
		return []byte("")
	}
	return
}
