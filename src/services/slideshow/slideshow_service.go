package slideshow

import (
	"main/dao"
	"main/services/database"
	"main/services/storage"
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

func (s *SlideShowService) GetShowById(id int) (SlideShowDTO, error) {
	return SlideShowDTO{}, nil
}

func (s *SlideShowService) GetSlideById(id int) (dao.Slide, error) {
	sl := dao.Slide{}
	slide, err := sl.GetById(s.Database.GetConnection(), id)
	if err != nil {
		return slide, err
	}
	return slide, nil
}

func (s *SlideShowService) SaveNewSlide(title string, file []byte) (dao.Slide, error) {
	sl := dao.Slide{Title: title}
	url, err := s.Storage.SaveBlob(file)
	if err != nil {
		return dao.Slide{}, err
	}
	sl.Url = url
	sl, err = sl.Save(s.Database.GetConnection())
	if err != nil {
		s.Storage.DeleteBlob(url)
		return dao.Slide{}, err
	}
	return sl, nil
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
