package structs

import (
	"github.com/tjones879/fake/util"
)

/*SavedPage TODO*/
type SavedPage struct {
	ID       string             `bson:"id"`
	Location string             `bson:"url"`
	Versions []util.FileStorage `bson:"versions"`
}

/*GetMostRecent TODO */
func (sp SavedPage) GetMostRecent() util.FileStorage {
	var recent util.FileStorage
	for _, fs := range sp.Versions {
		if fs.Date.After(recent.Date) {
			recent = fs
		}
	}
	return recent
}

/*GetByHash TODO */
func (sp SavedPage) GetByHash(hash uint64) util.FileStorage {
	return sp.Versions[0]
}
