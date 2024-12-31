package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	fmt.Println(s3Bucket)

	server := NewServer(3000)

	// server.AddMiddleware(middlewares.AuthMiddleware)

	server.AddHandler(handlers.MainEndpoint, handlers.MainEndpointHandler)
	server.AddHandler(handlers.UsersEndpoint, handlers.UsersEndpointHandler)

	server.Listen()
}
