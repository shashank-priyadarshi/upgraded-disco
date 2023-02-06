package main

import (
	"server/ghintegration"
	"server/server"
	"server/todos"
	"sync"
)

func main() {
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
