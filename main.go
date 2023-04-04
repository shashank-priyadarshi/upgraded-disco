package main

import (
	"crypto/rand"
	"encoding/base64"
	"os"
	"server/ghintegration"
	"server/server"
	"server/todos"
	"sync"
)

func main() {
	generateSigningKey()

	servers := []Server{
		&PrimaryServer{},
		&TodosServer{},
		&GHIntegrationServer{},
	}
	as := &AbstractServer{}
	as.StartServers(servers)
}

func generateSigningKey() {
	key := make([]byte, 1024)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	// Encode the key as base64 and store it in an environment variable
	encodedKey := base64.StdEncoding.EncodeToString(key)
	err = os.Setenv("SECRET_KEY", encodedKey)
	if err != nil {
		panic(err)
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

func (ts *TodosServer) StartServer() {
	todos.StartServer()
}

func (ghis *GHIntegrationServer) StartServer() {
	ghintegration.StartServer()
}
