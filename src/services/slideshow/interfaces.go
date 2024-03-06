package slideshow

import "main/dao"

type ShowService interface {
	Getter
	Saver
	Deleter
}

type Getter interface {
	GetShowById(id int) (*SlideShowData, error)
	GetSlideById(id int) (*dao.Slide, error)
}

type Saver interface {
	SaveNewSlideShow(showTitle string, slideTitle string, slide []byte) error
	SaveNewSlide(slideTitle string, slide []byte) error
}

type Deleter interface {
	DeleteSlideShow(id int) error
	DeleteSlide(id int) error
}
