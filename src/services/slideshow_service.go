package services

import (
	"log"
	"main/dao"
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
	*AzureSQLService
}

func (s *SlideShowService) GetByID(id int) (*SlideShowData, error) {
	log.Println(id)
	show := dao.NewSlideShow()
	show.ID = 1
	show.Title = "Hello"
	show.CreatedAt = time.Now()
	slide := dao.NewSlide("test-slide", "cat.jpg")
	_slides := []*dao.Slide{slide, slide}
	slides := []*SlideWithPosition{NewSlideWithPosition(_slides[0], 1)}
	return NewSlideShowData(show, slides), nil
}
