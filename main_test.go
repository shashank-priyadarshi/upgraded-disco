package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/mock"
)

type mockServer struct {
	mock.Mock
}

func (m *mockServer) StartServer() {
	m.Called()
}

func TestMain(t *testing.T) {
	var wg sync.WaitGroup

	server := new(mockServer)
	todos := new(mockServer)
	ghintegration := new(mockServer)

	server.On("StartServer").Once()
	todos.On("StartServer").Once()
	ghintegration.On("StartServer").Once()

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

	server.AssertExpectations(t)
	todos.AssertExpectations(t)
	ghintegration.AssertExpectations(t)
}
