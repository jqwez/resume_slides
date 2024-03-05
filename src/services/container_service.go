package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/joho/godotenv"
)

type AzureBlobService struct {
	*AzureBlobConfig
	*azblob.Client
}

func NewAzureBlobService(config *AzureBlobConfig) *AzureBlobService {
	service := &AzureBlobService{
		AzureBlobConfig: config,
	}
	client, err := service.Connect()
	if err != nil {
		log.Fatal("Failed to connect to Azure Blob Service")
	}
	service.Client = client
	err = service.CreateContainerIfNotExist(
		"https://127.0.0.1:10000/devstoreaccount1",
		"test-slides",
	)
	if err != nil {
		log.Println("Error Creating Container")
	}
	return service
}

type AzureBlobConfig struct {
	AccountName   string
	ContainerName string
	BaseURL       string
}

func AzureBlogConfigFromEnv() *AzureBlobConfig {
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

func (a *AzureBlobService) CreateContainerIfNotExist(baseURL string, containerName string) error {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	client, err := azblob.NewClient(baseURL, cred, nil)
	handleError(err)
	_, err = client.CreateContainer(
		context.TODO(),
		containerName,
		&azblob.CreateContainerOptions{
			Metadata: map[string]*string{"hello": to.Ptr("world")},
		})
	return err
}

func (a *AzureBlobService) Connect() (*azblob.Client, error) {
	containerURL := fmt.Sprintf("%s/%s/%s", a.BaseURL, a.AccountName, a.ContainerName)
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	containerClient, err := azblob.NewClient(containerURL, cred, nil)
	handleError(err)
	a.Client = containerClient
	log.Println("Connected to Azure Blob Service at:", a.BaseURL, a.AccountName, a.ContainerName)
	return containerClient, nil
}

func (a *AzureBlobService) GetBlobConnection() (*azblob.Client, error) {
	if a.Client != nil {
		return a.Client, nil
	}
	newConnection, err := a.Connect()
	a.Client = newConnection
	return a.Client, err
}

func (a *AzureBlobService) SaveBlob() (string, error) {
	bufferSize := 8 * 1024 * 1024
	blobName := "random" // will be random
	blobData := make([]byte, bufferSize)
	blobContentReader := bytes.NewReader(blobData)

	resp, err := a.Client.UploadStream(
		context.TODO(),
		"",
		blobName,
		blobContentReader,
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"hello": to.Ptr("world")},
		})
	handleError(err)
	log.Println(resp)
	return blobName, nil
}

func (a *AzureBlobService) DeleteBlobByName(name string) (bool, error) {
	resp, err := a.Client.DeleteBlob(context.TODO(), "", name, nil)
	if err != nil {
		return false, err
	}
	fmt.Println(resp)
	return true, nil
}

func (a *AzureBlobService) GetBlobByName(client *azblob.Client, name string) ([]byte, error) {
	downloadResponse, err := client.DownloadStream(
		context.TODO(),
		"",
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
		"",
		blobName,
		blob,
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"cat": to.Ptr("true")},
		})
	handleError(err)
	log.Println(resp)
	return blobName, nil
}
