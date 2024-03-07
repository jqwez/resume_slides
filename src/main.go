package main

import (
	"github.com/jqwez/resume_slides/controller"
	"github.com/jqwez/resume_slides/dao"
	"github.com/jqwez/resume_slides/services/database"
)

func main() {

	server := controller.NewServer()
	dbService := database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	dao.Migrate(dbService.GetConnection())
	server.Run()
}
