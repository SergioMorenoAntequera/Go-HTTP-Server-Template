package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port       int
	router     EndpointHandler
	middlwares []midleware
}

func NewServer(port int) *Server {
	return &Server{
		port:       port,
		router:     EndpointHandler{},
		middlwares: []midleware{},
	}
}

func (s *Server) AddMiddleware(newMidleware midleware) {
	s.middlwares = append(s.middlwares, newMidleware)
}

func (s *Server) AddRouter(endpoint string, endpointHandler EndpointHandler) {

	var functionality = func(w http.ResponseWriter, r *http.Request) {
		endpointHandler[r.Method](w, r)
	}

	s.router[endpoint] = functionality
}

func (s *Server) Listen() {

	for endpoint, endpointHandler := range s.router {

		var fullHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
			newW := w
			newR := r

			for _, midleware := range s.middlwares {
				canContinue, affectedW, affectedR := midleware(newW, newR)
				newW = affectedW
				newR = affectedR
				if !canContinue {
					return
				}
			}

			endpointHandler(w, r)
		}

		http.Handle(endpoint, fullHandler)
	}

	fmt.Printf("\n SERVER LISTENING IN: %d \n\n", s.port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	if err != nil {
		panic(err.Error())
	}
}
