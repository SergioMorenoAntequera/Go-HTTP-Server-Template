package handlers

import (
	"fmt"
	"net/http"
)

const UsersEndpoint = "/users"

var UsersEndpointHandler = EndpointHandler{
	"GET": func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s \n", r.Method, UsersEndpoint)
	},
}
