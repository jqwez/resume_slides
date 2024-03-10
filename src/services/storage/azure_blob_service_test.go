package storage

import (
	"context"
	"log"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

func TestMustAzureBlogConfigFromEnv(t *testing.T) {
	config := MustAzureBlogConfigFromEnv()
	if config.AccountName == "" || config.ContainerName == "" || config.BaseURL == "" {
		t.Fatal("Missing Env Value Azure Blog Config")
	}
}

func TestMustNewAzureBlobService(t *testing.T) {
	service := MustNewAzureBlobService(MustAzureBlogConfigFromEnv())
	_ = service
}

func TestCreateContainerIfNotExist(t *testing.T) {
	service := MustNewAzureBlobService(MustAzureBlogConfigFromEnv())
	err := service.CreateContainerIfNotExist("testy")
	if err != nil {
		t.Fatal(err)
	}
	service.Client.DeleteContainer(context.TODO(), "testy", &container.DeleteOptions{})
	log.Default()
}

func TestGetConnection(t *testing.T) {
	service := MustNewAzureBlobService(MustAzureBlogConfigFromEnv())
	_, err := service.GetConnection()
	if err != nil {
		t.Fatal("Failed ot get connection upon creation")
	}
	service.Client = nil
	_, err = service.GetConnection()
	if err != nil {
		t.Fatal("Ridded the connection and then failed to reconnect")
	}
}

func TestSaveAndDeleteBlob(t *testing.T) {
	service := MustNewAzureBlobService(MustAzureBlogConfigFromEnv())
	emptyFile := []byte{}
	fileName, err := service.SaveBlob(emptyFile)
	if err != nil {
		t.Fatal("Failed to Save Blob")
	}
	_, err = service.DeleteBlob(fileName)
	if err != nil {
		t.Fatal("Failed to Delete")
	}
}

func TestGetBlob(t *testing.T) {
	service := MustNewAzureBlobService(MustAzureBlogConfigFromEnv())
	emptyFile := []byte{}
	fileName, err := service.SaveBlob(emptyFile)
	if err != nil {
		t.Fatal("Failed to Save Blob")
	}
	blob, err := service.GetBlob(fileName)
	if err != nil {
		t.Fatal("Getting blob Error")
	}
	t.Log(blob)
	_, err = service.DeleteBlob(fileName)
	if err != nil {
		t.Fatal("Failed to Delete")
	}
	_, err = service.GetBlob(fileName)
	if err == nil {
		t.Fatal("Error should exist")
	}
}
