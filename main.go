package main

import (
	"server/server"
	"server/todos"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		server.StartServer()
	}()

	go func() {
		defer wg.Done()
		todos.StartServer()
	}()

	wg.Wait()
}
