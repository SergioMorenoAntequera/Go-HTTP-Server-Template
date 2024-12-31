package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

const UsersEndpoint = "/users/"

var UsersEndpointHandler = EndpointHandler{
	"GET": func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s \n", r.Method, UsersEndpoint)
		id := strings.Replace(r.URL.Path, UsersEndpoint, "", 1)

		if id == "" {
			fmt.Println("GET ALL USERS")
		}

		if id != "" {
			fmt.Println("GET USER WITH ID", id)
		}
	},
}
