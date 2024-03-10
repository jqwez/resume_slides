package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb" // Import MSSQL driver
	"github.com/joho/godotenv"
)

type AzureSQLService struct {
	*AzureSQLConfig
	conn *sql.DB
}

func NewAzureSQLService(config *AzureSQLConfig) *AzureSQLService {
	service := &AzureSQLService{
		AzureSQLConfig: config,
	}
	service.conn = service.connect()
	return service
}

func (a *AzureSQLService) GetConnection() *sql.DB {
	err := a.conn.Ping()
	if err != nil {
		a.conn = a.connect()
	}
	return a.conn
}

type AzureSQLConfig struct {
	Database string
	Password string
	User     string
	Server   string
}

func MustAzureSQLConfigFromEnv() *AzureSQLConfig {
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("error loading .env file")
		}
	}
	return &AzureSQLConfig{
		Database: os.Getenv("AZURE_DATABASE"),
		Password: os.Getenv("AZURE_PASSWORD"),
		User:     os.Getenv("AZURE_USER"),
		Server:   os.Getenv("AZURE_SERVER"),
	}
}

func (s *AzureSQLService) connect() *sql.DB {
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", s.Server, s.User, s.Password, s.Database)
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MSSQL at", s.Server)
	return db
}
