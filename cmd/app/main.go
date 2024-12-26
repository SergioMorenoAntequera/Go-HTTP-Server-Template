package main

import (
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
)

func main() {
	server := NewServer(3000)

	// server.AddMiddleware(middlewares.AuthMiddleware)

	server.AddHandler(handlers.MainEndpoint, handlers.MainEndpointHandler)
	server.AddHandler(handlers.UsersEndpoint, handlers.UsersEndpointHandler)

	server.Listen()
}
