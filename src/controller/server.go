package controller

import (
	"log"
	"net/http"
)

type Server struct {
	*Router
	port string
}

func NewServer() *Server {
	return &Server{
		Router: NewRouter(),
		port:   ":8000",
	}
}

func (s *Server) Run() {
	log.Println("Listening on port", s.port)
	http.ListenAndServe(s.port, nil)
}
