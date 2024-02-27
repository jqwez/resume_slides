package main

import (
	"log"
	"net/http"
	"main/router"
	"main/controller"
	"main/model"
)

func main() {
	db := controller.GetDatabaseConnection()
	model.Migrate(db)
	router.RegisterRoutes()
	port := ":8000"
	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}