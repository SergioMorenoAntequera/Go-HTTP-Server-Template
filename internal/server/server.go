package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/middlewares"
)

type Server struct {
	port       int
	middlwares []middlewares.Midleware
	mux        *http.ServeMux
}

func NewServer(port int) *Server {

	return &Server{
		port:       port,
		middlwares: []middlewares.Midleware{},
		mux:        http.NewServeMux(),
	}
}

func (s *Server) AddMiddleware(newMidleware middlewares.Midleware) {
	s.middlwares = append(s.middlwares, newMidleware)
}

func (s *Server) AddHandler(endpointHandler *handlers.EndpointHandler) {

	s.mux.HandleFunc(endpointHandler.Endpoint, func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s \n", r.Method, endpointHandler.Endpoint)

		handler, found := endpointHandler.Handler[r.Method]
		if !found {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

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

		handler(w, r)
	})
}

func (s *Server) Listen() {

	log.Printf("SERVER LISTENING IN: %d \n\n", s.port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.mux); err != nil {
		panic(err.Error())
	}
}
