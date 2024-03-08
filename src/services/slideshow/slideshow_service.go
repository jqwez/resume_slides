package slideshow

import (
	"log"

	"github.com/jqwez/resume_slides/dao"
	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/storage"
)

type SlideShowDTO struct {
	SlideShow *dao.SlideShow `json:"slideshow_data"`
	Slides    []*dao.Slide   `json:"slides"`
}

func NewSlideShowDTO(sw *dao.SlideShow, sl []*dao.Slide) *SlideShowDTO {
	return &SlideShowDTO{SlideShow: sw, Slides: sl}
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

func (s *SlideShowService) GetDb() database.DBService {
	return s.Database
}

func (s *SlideShowService) GetStore() storage.StorageService {
	return s.Storage
}

func (s *SlideShowService) GetAllShows() ([]SlideShowDTO, error) {
	sh := &dao.SlideShow{}
	slideshows, err := sh.GetAll(s.Database.GetConnection())

	if err != nil {
		return nil, err
	}
	dtos := []SlideShowDTO{}
	for _, show := range slideshows {
		dtos = append(dtos, SlideShowDTO{&show, nil})
	}
	return dtos, nil
}

func (s *SlideShowService) GetShowById(id int) (SlideShowDTO, error) {
	return SlideShowDTO{SlideShow: &dao.SlideShow{}, Slides: []*dao.Slide{&dao.Slide{Url: "cat.jpg"}, &dao.Slide{Url: "cat.jpg"}}}, nil
}

func (s *SlideShowService) GetAllSlides() (*SlideShowDTO, error) {
	sl := dao.Slide{}
	slides, err := sl.GetAll(s.Database.GetConnection())
	if err != nil {
		return &SlideShowDTO{}, err
	}
	return NewSlideShowDTO(nil, slides), nil

}

func (s *SlideShowService) GetSlideById(id int) (*dao.Slide, error) {
	sl := &dao.Slide{}
	slide, err := sl.GetById(s.Database.GetConnection(), id)
	if err != nil {
		return slide, err
	}
	return slide, nil
}

func (s *SlideShowService) SaveNewSlide(title string, file []byte) (*dao.Slide, error) {
	sl := &dao.Slide{Title: title}
	url, err := s.Storage.SaveBlob(file)
	if err != nil {
		return &dao.Slide{}, err
	}
	sl.Url = url
	sl, err = sl.Save(s.Database.GetConnection())
	if err != nil {
		s.Storage.DeleteBlob(url)
		return &dao.Slide{}, err
	}
	return sl, nil
}

func (s *SlideShowService) SaveNewSlideShow(showTitle string) (dao.SlideShow, error) {
	show := dao.NewSlideShow(showTitle)
	newSlideShow, err := show.Save(s.Database.GetConnection())
	if err != nil {
		return dao.SlideShow{}, err
	}
	return newSlideShow, nil
}

func (s *SlideShowService) DeleteSlideShow(id int) error {
	show := &dao.SlideShow{}
	_, err := show.DeleteById(s.Database.GetConnection(), id)
	if err != nil {
		log.Println("Error deleting show:", err)
		return err
	}
	return nil
}

func (s *SlideShowService) DeleteSlidePosition(id int) error {
	sl := &dao.SlidePosition{}
	slide, err := sl.GetById(s.Database.GetConnection(), id)
	if err != nil {
		log.Println("Could not find Slide")
		return err
	}
	_ = slide
	return nil
}
