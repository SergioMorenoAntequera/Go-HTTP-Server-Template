package main

import (
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/internal/models"
)

func translateToDBType(goType string) string {
	goToDBtype := map[string]string{
		"string": "VARCHAR(255)",
		"int":    "INT",
		"Time":   "TIMESTAMP",
	}

	if val, found := goToDBtype[goType]; found {
		return val
	}

	panic("Attribute type " + goType + " not found in DB parser")
}

type IModel interface {
	SetModelName()
}

func main() {
	// db.Connect()

	user := models.NewUser("Pedrito", 4)
	models.Create(user)

	/*
		http_server := server.NewServer(3000)
		server.AddMiddleware(middlewares.AuthMiddleware)
		http_server.AddHandler(handlers.MainEndpointHandler)
		http_server.AddHandler(handlers.UserEndpointHandler)
		http_server.AddHandler(handlers.UserEndpointHandlerOne)

		http_server.Listen()
	*/
}
