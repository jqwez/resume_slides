package controller

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/slideshow"
	"github.com/jqwez/resume_slides/services/storage"
)

type Router struct {
	slideshowService slideshow.ShowService
	staticDir        string
}

func NewRouter(db database.DBService, store storage.StorageService) *Router {
	router := &Router{
		slideshowService: slideshow.NewSlideShowService(db, store),
		staticDir:        staticLocationFinder(),
	}
	router.RegisterRoutes()
	return router
}

func (r *Router) RegisterRoutes() {
	http.HandleFunc("/", r.HandleGetHome)
	http.HandleFunc("/ws", HandleSocket)
	http.HandleFunc("/blob/{blobName}", r.HandleImageBlob)
	//	http.HandleFunc("/slideshow/{SlideShowId}", r.HandleGetSlideShowData)
	//	http.HandleFunc("/newslideshow", r.HandleNewSlideShow)
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
	client := r.slideshowService.GetStore()
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
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, "Error serving image", http.StatusInternalServerError)
		return
	}
}

/*
	func (r *Router) HandleGetSlideShowData(w http.ResponseWriter, req *http.Request) {
		SlideShowId := req.PathValue("SlideShowId")
		log.Println("Fetching slideshow: ", SlideShowId)

		slideshow_service := services.SlideShowService{AzureSQLService: r.AzureSQLService}
		ssd, err := slideshow_service.GetByID(1)
		if err != nil {
			log.Fatal(err)
		}

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

	func (r *Router) HandleNewSlideShow(w http.ResponseWriter, req *http.Request) {
		err := req.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			log.Println(err)
			return
		}
		file, _, err := req.FormFile("file")
		if err != nil {
			http.Error(w, "unable to get file from form", http.StatusBadRequest)
			log.Println(err)
			return
		}
		defer file.Close()

		fileContent, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Unable to read file", http.StatusInternalServerError)
			return
		}
		url, err := r.AzureBlobService.SaveBlob(fileContent)
		if err != nil {
			log.Println("Failed to save blob")
			http.Error(w, "Failed to save blob", http.StatusInternalServerError)
			return
		}
		err = r.SlideShowService.NewSlideShowEntry("test", url)
		if err != nil {
			log.Println("Failed to save slideshow")
			http.Error(w, "Failed to save slideshow", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "application/json")
	}
*/
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
