package database

import "database/sql"

type DBService interface {
	Connector
}

type Connector interface {
	connect() *sql.DB
	GetConnection() *sql.DB
}
