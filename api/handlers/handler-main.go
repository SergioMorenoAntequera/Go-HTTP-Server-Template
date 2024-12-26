package handlers

import (
	"fmt"
	"net/http"
)

const MainEndpoint = "/"

func MainEndpointHandler() EndpointHandler {
	methods := EndpointHandler{}

	methods["GET"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got %s request %s \n", MainEndpoint, r.Method)
		w.Write([]byte("Hello, World!"))
	}

	methods["POST"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("post %s request %s \n", MainEndpoint, r.Method)
	}

	return methods
}
