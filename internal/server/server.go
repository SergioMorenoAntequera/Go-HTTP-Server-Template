package server

import (
	"fmt"
	"net/http"

	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/handlers"
	"github.com/SergioMorenoAntequera/Go-HTTP-Server-Template/api/middlewares"
)

type Server struct {
	port       int
	middlwares []middlewares.Midleware
}

func NewServer(port int) *Server {
	return &Server{
		port:       port,
		middlwares: []middlewares.Midleware{},
	}
}

func (s *Server) AddMiddleware(newMidleware middlewares.Midleware) {
	s.middlwares = append(s.middlwares, newMidleware)
}

func (s *Server) AddHandler(endpoint string, endpointHandler handlers.EndpointHandler) {

	http.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {

		if handler, found := endpointHandler[r.Method]; found {

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
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}

	})
}

func (s *Server) Listen() {

	fmt.Printf("\n SERVER LISTENING IN: %d \n\n", s.port)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil); err != nil {
		panic(err.Error())
	}
}
