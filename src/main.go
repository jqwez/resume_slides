package main

import (
	"log"
	"main/controller"
	"main/model"
	"main/router"
	"net/http"
)

func main() {
	db := controller.GetDatabaseConnection()
	client, err := controller.GetContainerConnection()
	if err != nil {
		log.Fatal(err)
	}
	catText, err := controller.SaveCat(client)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(catText)
	model.Migrate(db)
	router.RegisterRoutes()
	port := ":8000"
	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
