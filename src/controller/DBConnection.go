package database

import (
	"database/sql"
	"os"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb" // Import MSSQL driver
	"github.com/joho/godotenv"
)

func DatabaseConnection() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env File")
	}
	database := os.Getenv("AZURE_DATABASE")
	password := os.Getenv("AZURE_PASSWORD")
	user := os.Getenv("AZURE_USER")
	server := os.Getenv("AZURE_SERVER")
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", server, user, password, database)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MSSQL at", server)
	return db
}