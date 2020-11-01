package main

import "fmt"

// StorageType defines available storage type
type StorageType int

const (
	// Mongo will use mongodb to storage
	Mongo StorageType = iota
	// Memory will use memory to storage for testing pourposes
	Memory
)

// Storage interface
type Storage interface {
	Find(string) ([]*Product, error)
}

// NewStorage ...
func NewStorage(storageType StorageType) (Storage, error) {
	var stg Storage
	var err error

	switch storageType {
	case Memory:
		stg = new(StorageMemory)

	case Mongo:
		uri := fmt.Sprintf(
			"mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
			goDotEnvVariable("USERNAME"),
			goDotEnvVariable("PASSWORD"),
			goDotEnvVariable("DBURL"),
			goDotEnvVariable("DBNAME"),
		)
		stg, err = NewStorageMongo(uri)
	}

	return stg, err
}
