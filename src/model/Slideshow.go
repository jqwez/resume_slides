package model

import (
	"time"
	"database/sql"
)

type SlideShow struct {
	ID					int				`json:"id" db: "id"`
	Title				string		`json:"title" db: "title"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`	
}

func NewSlideShow() *SlideShow {
	return &SlideShow{}
}

func (s *SlideShow) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slideshow', 'U') IS NULL
	CREATE TABLE slideshow (
		id INT PRIMARY KEY,
		title NVARCHAR(255),
		created_at DATETIME
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}