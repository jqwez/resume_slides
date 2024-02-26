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
	client, err := controller.GetContainerConnection()
	if err != nil {
		log.Fatal("could not establish container connection")
	}
	controller.SaveBlob(client)

	log.Println("ContainerClient", client)
	model.Migrate(db)
	router.RegisterRoutes()
	port := ":8000"
	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}