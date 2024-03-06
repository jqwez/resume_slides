package slideshow

import (
	"log"
	"main/dao"
	"main/services/database"
	"main/services/storage"
	"time"
)

type SlideWithPosition struct {
	*dao.Slide `json:"slide"`
	Position   int `json:"position"`
}

func NewSlideWithPosition(slide *dao.Slide, pos int) *SlideWithPosition {
	return &SlideWithPosition{Slide: slide, Position: pos}
}

type SlideShowData struct {
	*dao.SlideShow      `json:"slideshow_data"`
	SlidesWithPositions []*SlideWithPosition `json:"slides"`
}

func NewSlideShowData(slideshow *dao.SlideShow, slides []*SlideWithPosition) *SlideShowData {
	return &SlideShowData{SlideShow: slideshow, SlidesWithPositions: slides}
}

type SlideShowService struct {
	Database database.DBService
	Storage  storage.StorageService
}

func NewSlideShowService(db database.DBService, store storage.StorageService) *SlideShowService {
	return &SlideShowService{
		Database: db,
		Storage:  store,
	}
}

func (s *SlideShowService) GetShowById(id int) (*SlideShowData, error) {
	log.Println(id)
	show := dao.NewSlideShow("test")
	show.ID = 1
	show.Title = "Hello"
	show.CreatedAt = time.Now()
	slide := dao.NewSlide("test-slide", "cat.jpg")
	_slides := []*dao.Slide{slide, slide}
	slides := []*SlideWithPosition{NewSlideWithPosition(_slides[0], 1)}
	return NewSlideShowData(show, slides), nil
}

func (s *SlideShowService) GetSlideById(id int) (*dao.Slide, error) {
	return &dao.Slide{}, nil
}

func (s *SlideShowService) SaveNewSlide(title string, slide []byte) error {
	return nil
}

func (s *SlideShowService) SaveNewSlideShow(showTitle string, slideTitle string, slide []byte) error {
	return nil
}

func (s *SlideShowService) DeleteSlideShow(id int) error {
	return nil
}

func (s *SlideShowService) DeleteSlide(id int) error {
	return nil
}
