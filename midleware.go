package main

import "net/http"

type midleware = func(w http.ResponseWriter, r *http.Request) (bool, http.ResponseWriter, *http.Request)
