package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"server/ghintegration"
	"server/server"
	"strings"
	"sync"

	logger "github.com/rs/zerolog/log"
)

func main() {
	if strings.EqualFold("0", os.Getenv("SETUP")) {
		logger.Info().Msg("Running portfolio setup in dev environment")
	} else {
		logger.Info().Msg("Running portfolio setup in prod environment")
	}
	generateSigningKey()

	servers := []Server{
		&PrimaryServer{},
		&GHIntegrationServer{},
	}
	as := &AbstractServer{}
	as.StartServers(servers)
}

func generateSigningKey() {
	key := make([]byte, 1024)
	_, err := rand.Read(key)
	if err != nil {
		logger.Info().Err(err).Msg("Failed to generate signing key")
	}

	// Encode the key as base64 and store it in an environment variable
	encodedKey := base64.StdEncoding.EncodeToString(key)
	err = os.Setenv("SECRET_KEY", encodedKey)
	if err != nil {
		logger.Info().Err(err).Msg("Failed to set SECRET_KEY environment variable")
	}
}

func (as *AbstractServer) StartServers(servers []Server) {
	var wg sync.WaitGroup
	wg.Add(len(servers))

	for _, server := range servers {
		go func(s Server) {
			defer wg.Done()
			s.StartServer()
		}(server)
	}

	wg.Wait()
}

func (ps *PrimaryServer) StartServer() {
	// implementation for starting server1
	server.StartServer()
}

func (ghis *GHIntegrationServer) StartServer() {
	ghintegration.StartServer()
}
