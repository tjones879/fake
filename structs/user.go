package structs

// User is an authenticated user.
type User struct {
	Name  string   `bson:"name" json:"name"`
	Pages []string `bson:"pages" json:"pages"`
	Email string   `bson:"email" json:"email"`
	ID    string   `bson:"id" json:"sub"`
}
