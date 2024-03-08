package controller

import (
	"fmt"
	"net/http"
)

type StaticController struct {
	staticDir string
}

func NewStaticController(dir string) *StaticController {
	return &StaticController{staticDir: dir}
}

func (s *StaticController) RegisterRoutes() {
	s.ServeStatic()
}

func (s *StaticController) ServeStatic() {
	fmt.Println("Serving Static on: ", s.staticDir)
	fs := http.FileServer(http.Dir(staticLocationFinder()))
	http.Handle(fmt.Sprintf("%s/", s.staticDir), http.StripPrefix(fmt.Sprintf("%s/", s.staticDir), fs))
}
