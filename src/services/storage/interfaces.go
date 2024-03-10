package storage

import "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"

type StorageService interface {
	Connector
	Gettor
	Saver
	Deleter
}

type Connector interface {
	GetConnection() (*azblob.Client, error)
}

type Gettor interface {
	GetBlob(name string) ([]byte, error)
}

type Saver interface {
	SaveBlob(file []byte) (string, error)
}

type Deleter interface {
	DeleteBlob(name string) (bool, error)
}
