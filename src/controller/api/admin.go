package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jqwez/resume_slides/services/slideshow"
)

type AdminApiRequest struct {
	Sample string `json:"sample"`
}

type AdminApiResponse struct {
}

type AdminApi struct {
	slideshowService slideshow.ShowService
}

func NewAdminApi(service slideshow.ShowService) *AdminApi {
	return &AdminApi{slideshowService: service}
}

func adminPrefix() Prefix {
	return func(stem string) string {
		return fmt.Sprintf("/api/admin/%s", stem)
	}
}

func (a *AdminApi) RegisterRoutes() {
	pre := adminPrefix()
	http.HandleFunc(pre("slideshow/all"), a.HandleGetAllSlideShows)
	http.HandleFunc(pre("slideshow/new"), a.HandleCreateSlideShow)
	http.HandleFunc(pre("slideshow/delete"), a.HandleDeleteSlideShow)
	http.HandleFunc(pre("slideshow/edit"), a.HandleEditSlideShow)

	http.HandleFunc(pre("slide/all"), a.HandleGetAllSlides)
	http.HandleFunc(pre("slide/new"), a.HandleCreateSlide)
	http.HandleFunc(pre("slide/delete"), a.HandleDeleteSlide)
	http.HandleFunc(pre("slide/edit"), a.HandleEditSlide)
}

func (a *AdminApi) HandleCreateSlideShow(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	var requestBody AdminApiRequest
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error decoding JSON body", http.StatusBadRequest)
		return
	}
	title := requestBody.Sample
	_, err = a.slideshowService.SaveNewSlideShow(title)
	if err != nil {
		http.Error(w, "Failure saving slideshow on server", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Request processed Succesfully"))
}

func (a *AdminApi) HandleDeleteSlideShow(w http.ResponseWriter, r *http.Request) {
}

func (a *AdminApi) HandleEditSlideShow(w http.ResponseWriter, r *http.Request) {
}

func (a *AdminApi) HandleGetAllSlideShows(w http.ResponseWriter, r *http.Request) {
	slideshows, err := a.slideshowService.GetAllShows()
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not get slideshows", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(slideshows)
	if err != nil {
		log.Println(err)
		http.Error(w, "Could not marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func (a *AdminApi) HandleGetAllSlides(w http.ResponseWriter, r *http.Request) {
	allSlides, err := a.slideshowService.GetAllSlides()
	jsonData, err := json.Marshal(allSlides)
	_ = err
	log.Println(allSlides)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (a *AdminApi) HandleCreateSlide(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	multipartFile, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error retrieving file", http.StatusBadRequest)
		return
	}
	defer multipartFile.Close()
	file, err := io.ReadAll(multipartFile)
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	slide, _ := a.slideshowService.SaveNewSlide("test", file)
	_ = slide

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Request processed Succesfully"))
}

func (a *AdminApi) HandleDeleteSlide(w http.ResponseWriter, r *http.Request) {
}

func (a *AdminApi) HandleEditSlide(w http.ResponseWriter, r *http.Request) {
}
