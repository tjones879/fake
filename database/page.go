package database

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/tjones879/fake/structs"
	"gopkg.in/mgo.v2/bson"
)

var pages = "pages"

/*SavePage saves a page that has not been seen yet in mongo.*/
func SavePage(url string, file structs.FileStorage) (page structs.SavedPage, err error) {
	page = structs.SavedPage{
		ID:       uuid.NewV4().String(),
		Location: url,
		Versions: []string{
			file.UID,
		},
	}
	err = usingCollection(pages).Insert(page)
	return
}

/*GetSavedPage TODO */
func GetSavedPage(url string) (page structs.SavedPage) {
	err := usingCollection(pages).Find(bson.M{"url": url}).One(&page)
	if err != nil {
		fmt.Println("GetSavedPage:", err)
	}
	return page
}

/*UpdateSavedPage TODO */
func UpdateSavedPage(page structs.SavedPage) error {
	err := usingCollection(pages).Update(bson.M{"id": page.ID}, page)
	return err
}

/*IsPageSaved TODO */
func IsPageSaved(url string) bool {
	n, err := usingCollection(pages).Find(bson.M{"url": url}).Count()
	if err != nil {
		fmt.Println("IsPageSaved:", err)
	}
	return n > 0
}
