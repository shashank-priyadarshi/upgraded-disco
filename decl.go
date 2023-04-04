package main

type Server interface {
	StartServer()
}

type AbstractServer struct{}

type PrimaryServer struct {
	AbstractServer
}

type TodosServer struct {
	AbstractServer
}

type GHIntegrationServer struct {
	AbstractServer
}
