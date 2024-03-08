package controller

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jqwez/resume_slides/controller/api"
	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/slideshow"
	"github.com/jqwez/resume_slides/services/storage"
)

type Router struct {
	slideshowService slideshow.ShowService
}

func NewRouter(db database.DBService, store storage.StorageService) *Router {
	router := &Router{
		slideshowService: slideshow.NewSlideShowService(db, store),
	}
	router.RegisterRoutes()
	return router
}

func (r *Router) RegisterRoutes() {
	api.NewAdminApi(r.slideshowService).RegisterRoutes()
	api.NewApi(r.slideshowService).RegisterRoutes()
	NewStaticController("/static").RegisterRoutes()
	NewSocketHub().RegisterRoutes()
	http.HandleFunc("/", r.HandleGetHome)
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
	http.ServeFile(w, req, fmt.Sprintf("%s/index.html", staticLocationFinder()))
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
