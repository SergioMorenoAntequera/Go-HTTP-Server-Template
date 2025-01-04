package handlers

import (
	"net/http"
)

type Handler = map[string]http.HandlerFunc
type EndpointHandler struct {
	Endpoint string
	Handler  Handler
}

func NewEndpointHandler(endpoint string, handler Handler) *EndpointHandler {
	return &EndpointHandler{
		Endpoint: endpoint,
		Handler:  handler,
	}
}
