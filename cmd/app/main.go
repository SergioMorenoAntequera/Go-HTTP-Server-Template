package main

import (
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/internal/server"
)

func main() {

	http_server := server.NewServer(3000)
	// conn := db.Connect()
	// defer db.Disconnect(conn)

	// server.AddMiddleware(middlewares.AuthMiddleware)

	http_server.AddHandler(handlers.MainEndpointHandler)
	http_server.AddHandler(handlers.UserEndpointHandler)
	http_server.AddHandler(handlers.UserEndpointHandlerOne)

	http_server.Listen()

}
