package handlers

import (
	"net/http"
)

var MainEndpointHandler = NewEndpointHandler("/", Handler{
	"GET": func(w http.ResponseWriter, r *http.Request) {

	},
	"POST": func(w http.ResponseWriter, r *http.Request) {

	},
})
