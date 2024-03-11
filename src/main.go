package main

import (
	"github.com/jqwez/resume_slides/controller"
)

func main() {

	server := controller.NewServer()
	//dbService := database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	//dao.Migrate(dbService.GetConnection())
	server.Run()
}
