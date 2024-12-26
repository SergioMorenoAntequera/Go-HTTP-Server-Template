package middlewares

import "net/http"

func AuthMiddleware(w http.ResponseWriter, r *http.Request) (bool, http.ResponseWriter, *http.Request) {

	apikey := r.Header.Get("apikey")
	hasApiKey := apikey != ""

	if !hasApiKey {
		http.Error(w, "Unauthorized!", http.StatusUnauthorized)
		return false, nil, nil
	}

	// Has API Key check db

	return true, w, r
}
