package controller

import (
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	//RegisterRoutes()
}

/*
func TestServeHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	r.HandleGetHome(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Serve Home incorrect status code")
	}
}

/*
func TestServeImageBlob(t *testing.T) {
	client, _ := controller.GetContainerConnection()
	controller.SaveCat(client)
	req, err := http.NewRequest("GET", "/blob/cat.jpg", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	ServeImageBlob(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Log(rec.Body)
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Cat not served")
	}
}

func TestServeStatic(t *testing.T) {
	_, err := http.NewRequest("GET", "/static", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Static Serve Test Failed")
	}
}

func TestServeSlideShowData(t *testing.T) {
	req, err := http.NewRequest("GET", "/slideshow/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	HandleGetSlideShowData(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Log(rec.Body)
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Cat not served")
	}
	log.Println(rec.Body)
}
*/
