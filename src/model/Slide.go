package model

import (
	"time"
	"database/sql"
)

type Slide struct {
	ID					int				`json:"id" db: "id"`
	Title				string		`json:"title" db: "title"`
	Url					string		`json:"url" db: "url"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`	
	Slideshow		int				`json:"slideshow_id" db:"slideshow_id"`
	Position		int				`json:"position" db: "position"`
}

func NewSlide() *Slide {
	return &Slide{}
}

func (ss *Slide) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slides', 'U') IS NULL
	CREATE TABLE slides (
		id INT PRIMARY KEY,
		title NVARCHAR(255),
		url NVARCHAR(255),
		created_at DATETIME,
		slideshow_id INT,
		position INT,
		FOREIGN KEY (slideshow_id) REFERENCES slideshow(id)
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}
