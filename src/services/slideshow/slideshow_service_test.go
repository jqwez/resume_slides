package slideshow

import (
	"database/sql"
	"errors"
	"main/services/database"
	"main/services/storage"
	"os"
	"testing"
)

var dbService database.DBService
var storeService storage.StorageService
var service ShowService

func TestMain(m *testing.M) {
	dbService = database.NewAzureSQLService(database.MustAzureSQLConfigFromEnv())
	storeService = storage.MustNewAzureBlobService(storage.MustAzureBlogConfigFromEnv())
	service = NewSlideShowService(dbService, storeService)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestNewSlideShowService(t *testing.T) {
	slide, err := service.GetSlideById(10)
	handleGetError(err)
	t.Log(slide)
}

func TestGetDb(t *testing.T) {
	db := service.GetDb()
	if !errors.Is(db.GetConnection().Ping(), dbService.GetConnection().Ping()) {
		t.Fatal("Somehow not same")
	}
}

func TestGetStore(t *testing.T) {
	_ = service.GetStore()
}
func TestGetSlideById(t *testing.T) {
	_, err := service.GetSlideById(553434)
	if err == nil {
		t.Fatal("Should error if not found")
	}

}

func TestSaveNewSlide(t *testing.T) {
	slide, err := service.SaveNewSlide("test_slide", []byte{})
	handleGetError(err)
	_ = slide
	getSlide, err := service.GetSlideById(slide.ID)
	handleGetError(err)
	if slide.ID != getSlide.ID {
		t.Fatal("Mismatch Id")
	}
}

func handleGetError(err error) {
	if err != nil && isNoRowError(err) != true {
		t := testing.T{}
		t.Fatal(err)
	}
}

func isNoRowError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
