package slideshow

import (
	"database/sql"
	"errors"
	"os"
	"testing"

	"github.com/jqwez/resume_slides/services/database"
	"github.com/jqwez/resume_slides/services/storage"
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
	if slide.ID < 0 {
		t.Fatal("Invalid ID")
	}
}

func TestSaveNewSlideShow(t *testing.T) {
	show, err := service.SaveNewSlideShow("test_show")
	if err != nil {
		t.Fatal("SaveNewSlideShow Error: ", err)
	}
	if show.ID < 1 {
		t.Fatal("Invalid ShowId be greater than zero : ", show.ID)
	}
}

func TestDeleteSlideShow(t *testing.T) {
	show, err := service.SaveNewSlideShow("super_test")
	err = service.DeleteSlideShow(show.ID)
	if err != nil {
		t.Fatal("Failed to delete new slideshow")
	}
	err = service.DeleteSlideShow(-5)
	//if err == nil {
	//	t.Fatal("Failed to detect slideshow doesn't exist")
	//}
}

func TestDeleteSlidePosition(t *testing.T) {
	err := service.DeleteSlidePosition(-1)
	_ = err
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
