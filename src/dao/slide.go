package dao

import (
	"database/sql"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

type Slide struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url" db:"url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Position  int       `json:"position" db:"-"`
}

func NewSlide(title string, url string) *Slide {
	return &Slide{
		Title: title,
		Url:   url,
	}
}

func NewSlideFromRow(row *sql.Row) (Slide, error) {
	slide := Slide{}
	err := row.Scan(&slide.ID, &slide.Title, &slide.Url, &slide.CreatedAt)
	return slide, err
}

func NewSlidesFromRows(rows *sql.Rows) ([]Slide, error) {
	var slides []Slide
	for rows.Next() {
		var slide Slide
		if err := rows.Scan(&slide.ID, &slide.Title, &slide.Url, &slide.CreatedAt); err != nil {
			return nil, err
		}
		slides = append(slides, slide)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return slides, nil
}

func (s *Slide) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slides', 'U') IS NULL
	CREATE TABLE slides (
		id INT IDENTITY(1,1) PRIMARY KEY,
		title NVARCHAR(255),
		url NVARCHAR(255),
		created_at DATETIME,
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}

func (s *Slide) GetAll(db *sql.DB) ([]Slide, error) {
	sql := "SELECT * FROM slides"
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return NewSlidesFromRows(rows)
}

func (s *Slide) GetById(db *sql.DB, id int) (Slide, error) {
	query := "SELECT * FROM slides WHERE id=@Id"
	result := db.QueryRow(query,
		sql.Named("Id", id),
	)
	return NewSlideFromRow(result)
}

func (s *Slide) Save(db *sql.DB) (Slide, error) {
	s.CreatedAt = time.Now()
	query := `
	INSERT INTO slides (title, url, created_at)
	VALUES (@Title, @Url, @CreatedAt)`
	s.CreatedAt = time.Now()
	result := db.QueryRow(query,
		sql.Named("Title", s.Title),
		sql.Named("Url", s.Url),
		sql.Named("CreatedAt", s.CreatedAt),
	)
	return NewSlideFromRow(result)
}
