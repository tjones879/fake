package structs

// User is an authenticated user.
type User struct {
	Name        string          `bson:"name" json:"name"`
	Files       []FileReference `bson:"files" json:"files"`
	Email       string          `bson:"email" json:"email"`
	ID          string          `bson:"id" json:"sub"`
	Annotations []string        `bson:"annotations" json:"annotations"`
}

// FileReference holds a user's reference to a file
type FileReference struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

/*
 */
