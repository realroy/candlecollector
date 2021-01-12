package persistence

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// DiskPersistence ...
type DiskPersistence struct{}

// NewDiskPersistence ...
func NewDiskPersistence() *DiskPersistence {
	return &DiskPersistence{}
}

// Save ...
func (p DiskPersistence) Save(path string, data interface{}) error {
	log.Println("disk persistence is saving ...")
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	os.Mkdir("data", 0777)

	err = ioutil.WriteFile(path, byteData, 0644)
	if err != nil {
		log.Println("disk persistence saving error", err)

		return err
	}

	log.Println("disk persistence saving successfully")

	return nil
}

// Load ...
func (p DiskPersistence) Load(path string) interface{} {
	return nil
}
