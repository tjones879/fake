package structs

/*SavedPage TODO*/
type SavedPage struct {
	ID       string   `bson:"id"`
	Location string   `bson:"url"`
	Versions []string `bson:"versions"`
}

/*GetMostRecent TODO */
func (sp SavedPage) GetMostRecent() FileStorage {
	var recent FileStorage
	return recent
}
