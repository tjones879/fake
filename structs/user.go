package structs

// User is an authenticated user.
type User struct {
	Name        string   `bson:"name" json:"name"`
	FileIDs     []string `bson:"files" json:"files"`
	Email       string   `bson:"email" json:"email"`
	ID          string   `bson:"id" json:"sub"`
	Annotations []string `bson:"annotations" json:"annotations"`
}

/*
 */
