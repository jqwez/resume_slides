package controller

import (
	"database/sql"
	"os"
	"fmt"
	"log"
	_ "github.com/denisenkom/go-mssqldb" // Import MSSQL driver
	"github.com/joho/godotenv"
)

type AzureSQLConfig struct {
	Database	string
	Password	string
	User			string
	Server		string
}

func AzureSQLConfigFromEnv(prod bool) *AzureSQLConfig {
	if err := godotenv.Load(); err !=nil {
		if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("error loading .env file")
		}
	}
	return &AzureSQLConfig{
		Database: os.Getenv("AZURE_DATABASE"),
		Password: os.Getenv("AZURE_PASSWORD"),
		User: os.Getenv("AZURE_USER"),
		Server: os.Getenv("AZURE_SERVER"),
	}
}

func DatabaseConnection(c *AzureSQLConfig) *sql.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", c.Server, c.User, c.Password, c.Database)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MSSQL at", c.Server)
	return db
}

func GetDatabaseConnection() *sql.DB {
	return DatabaseConnection(AzureSQLConfigFromEnv(true))
}