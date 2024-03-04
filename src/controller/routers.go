package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/model"
	"main/services"
	"net/http"
	"os"
)

type Router struct {
	*services.AzureBlobService
	staticDir string
}

func NewRouter() *Router {
	router := &Router{
		AzureBlobService: services.NewAzureBlobService(services.AzureBlogConfigFromEnv()),
		staticDir:        staticLocationFinder(),
	}
	router.RegisterRoutes()
	return router
}

func (r *Router) RegisterRoutes() {
	http.HandleFunc("/", r.HandleGetHome)
	http.HandleFunc("/ws", HandleSocket)
	http.HandleFunc("/blob/{blobName}", r.HandleImageBlob)
	http.HandleFunc("/slideshow/{SlideShowId}", r.HandleGetSlideShowData)
	r.ServeStatic()
}

func (r *Router) ServeStatic() {
	fmt.Println("static Dir is : ", r.staticDir)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(r.staticDir))))
}

func (r *Router) HandleGetHome(w http.ResponseWriter, req *http.Request) {
	log.Println(req.URL)
	if req.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if req.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, req, fmt.Sprintf("%s/index.html", r.staticDir))
}

func (r *Router) HandleImageBlob(w http.ResponseWriter, req *http.Request) {
	blobName := req.PathValue("blobName")
	containerClient, _ := r.GetBlobConnection()
	blob, err := r.GetBlobByName(containerClient, blobName)
	if err != nil {
		http.Error(w, "Error getting blob", http.StatusInternalServerError)
		return
	}
	reader := bytes.NewReader(blob)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", blobName))
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
}

func (r *Router) HandleGetSlideShowData(w http.ResponseWriter, req *http.Request) {
	SlideShowId := req.PathValue("SlideShowId")
	log.Println("Fetching slideshow: ", SlideShowId)

	type SlideShowData struct {
		*model.SlideShow
		Slides [10]*model.Slide
	}
	slideshow := model.NewSlideShow()
	slideshow.ID = 1
	slideshow.Title = "test show"

	slide := model.NewSlide()
	slide.ID = 3
	slide.Title = "test slide"
	slide.Url = "http://localhost:8000/cat.jpg"
	slide.SlideShowId = 1
	slide.Position = 0

	ssd := &SlideShowData{
		SlideShow: slideshow,
	}
	ssd.Slides[0] = slide

	jsonData, err := json.Marshal(ssd)
	if err != nil {
		http.Error(w, "Error Serving SlideShow", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func staticLocationFinder() string {
	serverDir := "/app/static"
	_, err := os.Stat(serverDir)
	if err == nil {
		return serverDir
	}
	log.Println("no app/static directory ----> try ../static")
	localDir := "../static"
	_, err = os.Stat(localDir)
	if err == nil {
		return localDir
	}
	log.Println("no ../static directory ----> try ../../static")
	testDir := "../../static"
	_, err = os.Stat(testDir)
	if err == nil {
		return testDir
	}
	log.Fatal("no static dir found")
	return ""
}
