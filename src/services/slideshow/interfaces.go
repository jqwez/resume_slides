package slideshow

import (
	"github.com/jqwez/resume_slides/dao"
	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/storage"
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
	GetSlideById(id int) (*dao.Slide, error)
	GetAllShows() ([]SlideShowDTO, error)
	GetAllSlides() (*SlideShowDTO, error)
}

type Saver interface {
	SaveNewSlideShow(showTitle string) (dao.SlideShow, error)
	SaveNewSlide(slideTitle string, slide []byte) (*dao.Slide, error)
}

type Deleter interface {
	DeleteSlideShow(id int) error
	DeleteSlidePosition(id int) error
}
