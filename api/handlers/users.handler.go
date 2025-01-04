package handlers

import (
	"net/http"
)

var UserEndpointHandler = NewEndpointHandler("/users", Handler{
	"GET": func(w http.ResponseWriter, r *http.Request) {

	},
})

var UserEndpointHandlerOne = NewEndpointHandler("/users/{id}", Handler{
	"GET": func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // Extract the {id} parameter
		w.Write([]byte("Post ID: " + id))
	},
})
