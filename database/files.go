package database

import (
	"github.com/tjones879/fake/structs"
	"gopkg.in/mgo.v2/bson"
)

var files = "files"

// GetFileByID TODO
func GetFileByID(uid string) (f structs.FileStorage, err error) {
	err = usingCollection(files).Find(bson.M{"id": uid}).One(&f)
	return
}

// SaveFile TODO
func SaveFile(f structs.FileStorage) error {
	return usingCollection(files).Insert(f)
}
