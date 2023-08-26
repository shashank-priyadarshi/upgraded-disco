package main

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate # once per package

//counterfeiter:generate . Server # for every interface
type Server interface {
	StartServer()
}

type AbstractServer struct{}

type PrimaryServer struct {
	AbstractServer
}

type GHIntegrationServer struct {
	AbstractServer
}
