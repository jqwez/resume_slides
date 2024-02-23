package routes

import (
	"log"
	"net/http"
)


func RegisterRoutes() {
	http.HandleFunc("/", ServeHome)
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
	http.ServeFile(w, r, "/app/static/index.html")
}

func ServeStatic() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/app/static"))))
}