package main

import "net/http"

func AuthMiddleware(w http.ResponseWriter, r *http.Request) (bool, http.ResponseWriter, *http.Request) {

	apikey := r.Header.Get("apikey")
	hasApiKey := apikey != ""

	if !hasApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return false, nil, nil
	}

	// Has API Key check db

	return true, w, r
}
