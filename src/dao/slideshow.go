package dao

import (
	"database/sql"
	"time"
)

type SlideShow struct {
	ID        int       `json:"id" db: "id"`
	Title     string    `json:"title" db: "title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type SlideShowData struct {
	*SlideShow
	Slides []*Slide
}

func NewSlideShow() *SlideShow {
	return &SlideShow{}
}

func NewSlideShowData(ss *SlideShow, sls []*Slide) *SlideShowData {
	return &SlideShowData{
		SlideShow: ss,
		Slides:    sls,
	}
}

func (s *SlideShow) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slideshows', 'U') IS NULL
	CREATE TABLE slideshows (
		id INT PRIMARY KEY,
		title NVARCHAR(255),
		created_at DATETIME
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}
