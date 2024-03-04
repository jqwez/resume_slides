package main

import (
	"main/controller"
	"main/model"
	"main/services"
)

func main() {
	db := services.GetDatabaseConnection()
	model.Migrate(db)

	server := controller.NewServer()

	server.Run()
}
