package dao

import (
	"database/sql"
	"log"
)

type SlidePosition struct {
	SlideShowID int `json:"slideshow_id" db: "slideshow_id"`
	SlideID     int `json:"slide_id" db: "slide_id"`
	Position    int `json:"position" db: "position"`
}

func NewSlidePosition(showID int, slideID int, position int) *SlidePosition {
	return &SlidePosition{
		SlideShowID: showID,
		SlideID:     slideID,
		Position:    position,
	}
}
func NewSlidePositionFromRow(row *sql.Row) (SlidePosition, error) {
	pos := SlidePosition{}
	err := row.Scan(&pos.SlideShowID, &pos.SlideID, &pos.Position)
	return pos, err
}

func NewSlidePositionsFromRows(rows *sql.Rows) ([]SlidePosition, error) {
	var positions []SlidePosition
	for rows.Next() {
		var pos SlidePosition
		if err := rows.Scan(&pos.SlideShowID, &pos.SlideID, &pos.Position); err != nil {
			return nil, err
		}
		positions = append(positions, pos)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return positions, nil
}

func (s *SlidePosition) Migrate(db *sql.DB) error {
	migrationSQL := `
	IF OBJECT_ID('slide_positions', 'U') IS NULL
	CREATE TABLE slide_positions (
		slideshow_id INT,
		slide_id INT,
		position INT,
	)
		FOREIGN KEY (slideshow_id) REFERENCES slideshows(id),
		FOREIGN KEY (slide_id) REFERENCES slides(id)
	);
	`
	_, err := db.Exec(migrationSQL)
	return err
}

func (s *SlidePosition) GetAll(db *sql.DB) ([]SlidePosition, error) {
	query := "SELECT * FROM slide_positions"
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	return NewSlidePositionsFromRows(rows)
}
