package main

import (
	"fmt"
	"net/http"
)

const MainEndpoint = "/"

func MainEndpointHandler() EndpointHandler {
	methods := make(EndpointHandler)

	methods["GET"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got %s request %s \n", MainEndpoint, r.Method)
	}

	methods["POST"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("post %s request %s \n", MainEndpoint, r.Method)
	}

	return methods
}
