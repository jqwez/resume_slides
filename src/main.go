package main

import (
	"main/controller"
	"main/dao"
	"main/services/database"
)

func main() {

	server := controller.NewServer()
	dbService := database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	dao.Migrate(dbService.GetConnection())
	server.Run()
}
