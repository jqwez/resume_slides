package router

import (
	"fmt"
	"log"
	"net/http"
	"io"
	"os"
	"bytes"

	"main/controller"
)

var StaticDir = staticLocationFinder()
var clientConn, err = controller.GetContainerConnection()

func RegisterRoutes() {
	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", SocketHandler)
	http.HandleFunc("/blob", ServeImageBlob)
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
	blob, err := controller.GetBlobByName(clientConn, "random")
	if err != nil {
		http.Error(w, "Error getting blob", http.StatusInternalServerError)	
		return
	}
	reader := bytes.NewReader(blob)
	w.Header().Set("Content-Type", "text/plain")
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
