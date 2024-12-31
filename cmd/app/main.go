package main

import (
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/internal/db"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/internal/server"
)

func main() {
	http_server := server.NewServer(3000)

	// server.AddMiddleware(middlewares.AuthMiddleware)

	http_server.AddHandler(handlers.MainEndpoint, handlers.MainEndpointHandler)
	http_server.AddHandler(handlers.UsersEndpoint, handlers.UsersEndpointHandler)

	db.Connect()
	http_server.Listen()
}
