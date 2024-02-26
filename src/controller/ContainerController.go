package controller

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"io"
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/joho/godotenv"
	//"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
)

type AzureBlobConfig struct {
	AccountName			string
	ContainerName		string
}

func AzureBlogConfigFromEnv() *AzureBlobConfig {
	if err := godotenv.Load(); err !=nil {
		if err := godotenv.Load("../.env"); err != nil {
		log.Fatal("error loading .env file")
		}
	}
	return &AzureBlobConfig{
		AccountName: os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"), 
		ContainerName: os.Getenv("AZURE_APP_CONTAINER_NAME"),
	}
}

func CreateContainerIfNotExist(containerName string) error {
	return nil
}

func ContainerConnection(c *AzureBlobConfig) (*azblob.Client, error) {
	containerURL := fmt.Sprintf("https://127.0.0.1:10000/%s/%s", c.AccountName, c.ContainerName)
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)
	containerClient, err := azblob.NewClient(containerURL, cred, nil)
	handleError(err) 
	return containerClient, nil
}

func GetContainerConnection() (*azblob.Client, error) {
	return ContainerConnection(AzureBlogConfigFromEnv())
}

func SaveBlob(client *azblob.Client) string {
	bufferSize := 8 * 1024 * 1024
	blobName := "random" // will be random
	blobData := make([]byte, bufferSize)
	blobContentReader := bytes.NewReader(blobData)

	resp, err := client.UploadStream(
		context.TODO(),
		"",
		blobName,
		blobContentReader,
		&azblob.UploadStreamOptions{
			Metadata: map[string]*string{"hello": to.Ptr("world")},
		},)
		handleError(err)
		log.Println(resp)
		return blobName
}

func DeleteBlobByName(client *azblob.Client, name string) error {
	resp, err := client.DeleteBlob(context.TODO(), "", name, nil)
	if err != nil {
		return err
	}
	fmt.Println(resp)
	return nil
}

func GetBlobByName(client *azblob.Client, name string) []byte {
	downloadResponse, err := client.DownloadStream(
		context.TODO(),
		"", 
		name,
		nil,
	)
	handleError(err)
	actualBlobData, err := io.ReadAll(downloadResponse.Body)
	handleError(err)
	return actualBlobData
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}