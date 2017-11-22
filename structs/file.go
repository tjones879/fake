package structs

import (
	"os"
	"time"
)

// FileStorage TODO
type FileStorage struct {
	UID       string    `bson:"id"`
	Hash      uint64    `bson:"hash"`
	Directory string    `bson:"dir"`
	Name      string    `bson:"name"`
	Contents  string    `bson:"-"`
	Date      time.Time `bson:"timestamp"`
}

// FileExists TODO
func (f FileStorage) FileExists() bool {
	_, err := os.Stat(f.Directory + "/" + f.Name)
	return !os.IsNotExist(err)
}
