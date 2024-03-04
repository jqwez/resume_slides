package dao

import "database/sql"

type DataAccessor interface {
	Migrator
	Gettor
}

type Gettor interface {
	GetAll(db *sql.DB) ([]interface{}, error)
}
