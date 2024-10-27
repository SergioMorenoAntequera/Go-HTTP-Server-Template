package main

func main() {
	server := NewServer(3000)

	server.AddMiddleware(AuthMiddleware)

	server.AddRouter(MainEndpoint, MainEndpointHandler())

	server.Listen()
}
