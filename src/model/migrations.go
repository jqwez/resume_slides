package model

import (
	"database/sql"
)

func Migrate(db *sql.DB) {
	s := NewSlideShow()
	sl := NewSlide()
	s.Migrate(db)
	sl.Migrate(db)
}