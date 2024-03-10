package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jqwez/resume_slides/services/slideshow"
)

type ApiRequest struct {
	Sample string `json:"sample"`
}

type ApiResponse struct {
}

type Api struct {
	slideshowService slideshow.ShowService
}

func NewApi(service slideshow.ShowService) *Api {
	return &Api{slideshowService: service}
}

func apiPrefix() Prefix {
	return func(stem string) string {
		return fmt.Sprintf("/api/%s", stem)
	}
}

func (a *Api) RegisterRoutes() {
	pre := apiPrefix()
	http.HandleFunc(pre("blob/{fileName}"), a.HandleGetBlob)
	http.HandleFunc(pre("slideshow"), a.HandleGetSlideShow)
}

func (a *Api) HandleGetBlob(w http.ResponseWriter, r *http.Request) {
	blobName := r.PathValue("fileName")
	client := a.slideshowService.GetStore()
	blob, err := client.GetBlob(blobName)
	if err != nil {
		http.Error(w, "Error getting blob", http.StatusInternalServerError)
		return
	}
	reader := bytes.NewReader(blob)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", blobName))
	w.Header().Set("Cache-Control", "public, max-age=100000")
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
	log.Println("/api/blob/", blobName)
}

func (a *Api) HandleGetSlideShow(w http.ResponseWriter, r *http.Request) {
	data, _ := a.slideshowService.GetShowById(1)
	jsonData, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Write(jsonData)
}
