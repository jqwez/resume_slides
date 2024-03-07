package slideshow

import (
	"main/dao"
	"main/services/database"
	"main/services/storage"
)

type ShowService interface {
	ServiceHolder
	Getter
	Saver
	Deleter
}

type ServiceHolder interface {
	GetDb() database.DBService
	GetStore() storage.StorageService
}

type Getter interface {
	GetShowById(id int) (SlideShowDTO, error)
	GetSlideById(id int) (dao.Slide, error)
}

type Saver interface {
	SaveNewSlideShow(showTitle string, slideTitle string, slide []byte) error
	SaveNewSlide(slideTitle string, slide []byte) (dao.Slide, error)
}

type Deleter interface {
	DeleteSlideShow(id int) error
	DeleteSlide(id int) error
}
