package main

import (
	"server/server"
	"server/todos"
)

func main() {
	server.StartServer()
	todos.StartServer()
}
