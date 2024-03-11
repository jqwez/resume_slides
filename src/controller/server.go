package controller

import (
	"log"
	"net/http"

	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/storage"
)

type Server struct {
	Router *Router
	port   string
}

func NewServer() *Server {
	dbService := database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	blobService := storage.MustNewAzureBlobService(storage.MustAzureBlogConfigFromEnv())
	return &Server{
		Router: NewRouter(dbService, blobService),
		port:   ":80",
	}
}

func (s *Server) Run() {
	log.Println("Listening on port", s.port)
	http.ListenAndServe(s.port, nil)
}
