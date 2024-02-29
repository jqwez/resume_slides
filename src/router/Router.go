package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"main/controller"
	"main/model"
)

var StaticDir = staticLocationFinder()
var clientConn, _ = controller.GetContainerConnection()

func RegisterRoutes() {
	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", SocketHandler)
	http.HandleFunc("/blob/{blobName}", ServeImageBlob)
	http.HandleFunc("/slideshow/{SlideShowId}", ServeSlideShowData)
	ServeStatic()
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, fmt.Sprintf("%s/index.html", StaticDir))
}

func ServeImageBlob(w http.ResponseWriter, r *http.Request) {
	blobName := r.PathValue("blobName")
	blob, err := controller.GetBlobByName(clientConn, blobName)
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

func ServeStatic() {
	fmt.Println("static Dir is : ", StaticDir)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(StaticDir))))
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

func ServeSlideShowData(w http.ResponseWriter, r *http.Request) {
	SlideShowId := r.PathValue("SlideShowId")
	log.Println(SlideShowId)

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
