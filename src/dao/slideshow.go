package dao

import (
	"database/sql"
	"time"
)

type SlideShow struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func NewSlideShow(title string) *SlideShow {
	return &SlideShow{Title: title}
}

func NewSlideShowFromRow(row *sql.Row) (SlideShow, error) {
	show := SlideShow{}
	err := row.Scan(&show.ID, &show.Title, &show.CreatedAt)
	return show, err
}

func (s *SlideShow) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slideshows', 'U') IS NULL
	CREATE TABLE slideshows (
		id INT IDENTITY(1, 1) PRIMARY KEY,
		title NVARCHAR(255),
		created_at DATETIME
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}

func (*SlideShow) GetById(db *sql.DB, id int) (SlideShow, error) {
	query := "SELECT * FROM slideshows WHERE id=@Id"
	result := db.QueryRow(query,
		sql.Named("Id", id),
	)
	return NewSlideShowFromRow(result)

}

func (s *SlideShow) Save(db *sql.DB) (SlideShow, error) {
	query := `
	INSERT INTO slideshows (title, created_at)
	VALUES (@Title, @CreatedAt);
	SELECT SCOPE_IDENTITY();
	`
	s.CreatedAt = time.Now()
	var id int64
	err := db.QueryRow(query,
		sql.Named("Title", s.Title),
		sql.Named("CreatedAt", s.CreatedAt),
	).Scan(&id)
	if err != nil {
		return SlideShow{}, err
	}
	newSlideShow, err := s.GetById(db, int(id))
	if err != nil {
		return SlideShow{}, err
	}
	return newSlideShow, nil
}

func (s *SlideShow) DeleteById(db *sql.DB, id int) (bool, error) {
	query := `
	DELETE FROM slideshows
	WHERE id=@Id;
	`
	_, err := db.Exec(query,
		sql.Named("Id", id),
	)
	if err != nil {
		return false, err
	}
	return true, nil
}
