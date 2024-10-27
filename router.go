package main

import (
	"net/http"
)

type EndpointHandler = map[string]http.HandlerFunc
