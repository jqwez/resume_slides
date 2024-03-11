package storage

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"
)

type AzureBlobConfig struct {
	AccountName   string
	ContainerName string
	BaseURL       string
}

func MustAzureBlogConfigFromEnv() *AzureBlobConfig {
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../.env"); err != nil {
			log.Fatal("error loading .env file")
		}
	}
	return &AzureBlobConfig{
		AccountName:   os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"),
		ContainerName: os.Getenv("AZURE_APP_CONTAINER_NAME"),
		BaseURL:       os.Getenv("AZURE_BASE_URL"),
	}
}

type AzureBlobService struct {
	Config *AzureBlobConfig
	Client *azblob.Client
}

func MustNewAzureBlobService(config *AzureBlobConfig) *AzureBlobService {
	service := &AzureBlobService{
		Config: config,
	}
	client, err := service.connect()
	if err != nil {
		log.Fatal("Failed to connect to Azure Blob Service")
	}
	service.Client = client
	err = service.CreateContainerIfNotExist(
		"test-slides",
	)
	if err != nil {
		log.Println("Error Creating Container")
	}
	return service
}

func (a *AzureBlobService) CreateContainerIfNotExist(containerName string) error {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	// fmt.Sprintf("%s/%s", a.Config.BaseURL, a.Config.AccountName)
	url := "https://resumeslidesblob.blob.core.windows.net/"
	client, err := azblob.NewClient(url, cred, nil)
	handleError(err)
	_, err = client.CreateContainer(
		context.Background(),
		containerName,
		&azblob.CreateContainerOptions{
			Metadata: map[string]*string{"hello": to.Ptr("world")},
		})
	return err
}

func (a *AzureBlobService) connect() (*azblob.Client, error) {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	// url := fmt.Sprintf("%s/%s", a.Config.BaseURL, a.Config.AccountName)
	url := "https://resumeslidesblob.blob.core.windows.net/"
	client, err := azblob.NewClient(url, cred, nil)
	handleError(err)
	a.Client = client
	log.Println("Connected to Azure Blob Service at:", url)
	// SaveCat(a.Client)
	return a.Client, nil
}

func (a *AzureBlobService) GetConnection() (*azblob.Client, error) {
	if a.Client != nil {
		return a.Client, nil
	}
	newConnection, err := a.connect()
	a.Client = newConnection
	return a.Client, err
}

func makeFileName() (string, error) {
	numBytes := 64 / 2
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	randomFileName := hex.EncodeToString(randomBytes)
	return randomFileName, nil
}

func (a *AzureBlobService) SaveBlob(file []byte) (string, error) {
	blobName, err := makeFileName()
	handleError(err)
	blobContentReader := bytes.NewReader(file)

	resp, err := a.Client.UploadStream(
		context.TODO(),
		a.Config.ContainerName,
		blobName,
		blobContentReader,
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"hello": to.Ptr("world")},
		})
	handleError(err)
	log.Println(resp)
	return blobName, nil
}

func (a *AzureBlobService) DeleteBlob(name string) (bool, error) {
	resp, err := a.Client.DeleteBlob(context.TODO(), a.Config.ContainerName, name, nil)
	if err != nil {
		return false, err
	}
	fmt.Println(resp)
	return true, nil
}

func (a *AzureBlobService) GetBlob(name string) ([]byte, error) {
	downloadResponse, err := a.Client.DownloadStream(
		context.TODO(),
		a.Config.ContainerName,
		name,
		nil,
	)
	if err != nil {
		return nil, err
	}
	actualBlobData, err := io.ReadAll(downloadResponse.Body)
	if err != nil {
		return nil, err
	}
	return actualBlobData, nil
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func SaveCat(client *azblob.Client) (string, error) {
	blob, err := os.Open("../static/cat.jpg")
	if err != nil {
		blob, err = os.Open("../../static/cat.jpg")
	}
	if err != nil {
		log.Fatal("no cat found to upload")
	}
	blobName := "cat.jpg"
	resp, err := client.UploadStream(
		context.TODO(),
		"test-slides",
		blobName,
		blob,
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"cat": to.Ptr("true")},
		})
	handleError(err)
	log.Println(resp)
	return blobName, nil
}
