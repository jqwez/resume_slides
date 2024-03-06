package storage

type StorageService interface {
	//Connector
	//Gettor
	//Saver
	//Deleter
}

type Connector interface {
	GetConnection()
}

type Gettor interface {
	GetBlob(name string)
}

type Saver interface {
	SaveBlob(file []byte)
}

type Deleter interface {
	DeleteBlob(name string)
}
