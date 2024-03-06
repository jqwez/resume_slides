package controller

import (
	"log"
	"main/services/database"
	"main/services/storage"
	"net/http"
)

type Server struct {
	Router *Router
	port   string
}

func NewServer() *Server {
	dbService := database.NewAzureSQLService(database.AzureSQLConfigFromEnv())
	blobService := storage.NewAzureBlobService(storage.AzureBlogConfigFromEnv())
	return &Server{
		Router: NewRouter(dbService, blobService),
		port:   ":8000",
	}
}

func (s *Server) Run() {
	log.Println("Listening on port", s.port)
	http.ListenAndServe(s.port, nil)
}
