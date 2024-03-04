package dao

import (
	"database/sql"
)

type Migrator interface {
	Migrate(db *sql.DB) error
}

func Migrate(db *sql.DB) {
	migrations := []Migrator{
		&Slide{},
		&SlideShow{},
		&SlidePosition{},
	}
	for _, dao := range migrations {
		dao.Migrate(db)
	}
}
