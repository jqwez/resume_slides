package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterRoutes(t *testing.T) {
	RegisterRoutes()
}

func TestServeHome(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	ServeHome(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Serve Home incorrect status code")
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
