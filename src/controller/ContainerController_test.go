package controller

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
}


func TestContainerConnection(t *testing.T) {
	client, err := GetContainerConnection()
	assert.NotNil(t, client)
	assert.Nil(t, err)
}
func TestSaveBlob(t *testing.T) {
	client, _ := GetContainerConnection()
	location := SaveBlob(client)
	if location != "random" {
		t.Fatal("not random")
	}

}
func TestDeleteBlobByName(t *testing.T) {
	client, _ := GetContainerConnection()
	err := DeleteBlobByName(client, "testblob")
	if err != nil {
		t.Fatal("Delete blob failed: ", err)
	}
}

func TestGetBlobByName(t *testing.T) {
	client, _ := GetContainerConnection()
	newBlobName := SaveBlob(client)
	blob := GetBlobByName(client, newBlobName)
	if blob == nil {
		t.Fatal("no blob")
	}
}
