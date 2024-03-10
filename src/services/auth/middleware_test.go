package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequireAdmin(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()
	handleTestRoute(rec, req)
	if status := rec.Code; status != http.StatusOK {
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Serve incorrect status code")
	}
	t.Log("Expect Good Auth :", rec.Body)
	rec = httptest.NewRecorder()
	req.Header.Add("Authorization", "Bearer InvalidToken")
	handler := RequireAdmin(handleTestRoute)
	handler(rec, req)
	if status := rec.Code; status != http.StatusUnauthorized {
		t.Log(rec.Code)
		t.Log(status)
		t.Errorf("Serve Home incorrect status code")
	}
	t.Log("Expect Unauthorized :", rec.Body)
}

func handleTestRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good Auth"))
}
