package main

import (
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/middlewares"
)

func main() {
	server := NewServer(3000)

	server.AddMiddleware(middlewares.AuthMiddleware)

	server.AddRouter(handlers.MainEndpoint, handlers.MainEndpointHandler())

	server.Listen()
}
