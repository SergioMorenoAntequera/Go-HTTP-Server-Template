package handlers

import (
	"fmt"
	"net/http"
)

const TestEndpoint = "/test"

func TestEndpointHandler() EndpointHandler {
	methods := make(EndpointHandler)

	methods["GET"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("got %s request %s \n", TestEndpoint, r.Method)
	}

	methods["POST"] = func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("post %s request %s \n", TestEndpoint, r.Method)
	}

	return methods
}
