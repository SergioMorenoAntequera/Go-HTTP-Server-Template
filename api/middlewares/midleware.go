package middlewares

import "net/http"

type Midleware = func(w http.ResponseWriter, r *http.Request) (bool, http.ResponseWriter, *http.Request)
