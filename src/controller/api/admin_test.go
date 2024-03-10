package api

import (
	"os"
	"testing"

	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/slideshow"
	"github.com/jqwez/resume_slides/services/storage"
)

var slideshowService slideshow.ShowService

func TestMain(m *testing.M) {
	dbService := database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	storeService := storage.MustNewAzureBlobService(storage.MustAzureBlogConfigFromEnv())
	slideshowService = slideshow.NewSlideShowService(dbService, storeService)
	exitCode := m.Run()
	os.Exit(exitCode)
}
func TestNewAdminApi(t *testing.T) {
	_ = NewAdminApi(slideshowService)
}

func TestRegisterAdminRoutes(t *testing.T) {
	api := NewAdminApi(slideshowService)
	api.RegisterRoutes()
}

func TestHandleNewSlideShow(t *testing.T) {

}

func TestHandleGetAllSlideShows(t *testing.T) {
}
