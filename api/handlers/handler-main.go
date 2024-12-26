package handlers

import (
	"fmt"
	"net/http"
)

const MainEndpoint = "/main"

var MainEndpointHandler = EndpointHandler{
	"GET": func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s - %s \n", r.Method, MainEndpoint)
	},
}
