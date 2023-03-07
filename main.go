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
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		server.StartServer()
	}()

	go func() {
		defer wg.Done()
		todos.StartServer()
	}()

	go func() {
		defer wg.Done()
		ghintegration.StartServer()
	}()

	wg.Wait()
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
