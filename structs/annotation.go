package structs

import (
	"github.com/satori/go.uuid"
	"time"
)

/*
type User struct {
	Name  string   `bson:"name" json:"name"`
	Pages []string `bson:"pages" json:"pages"`
	Email string   `bson:"email" json:"email"`
	ID    string   `bson:"id" json:"sub"`
}
*/

type (
	// Metadata stores info about a RESTful API.
	Metadata struct {
		Message   string    `json:"message"`
		Endpoints []APILink `json:"links"`
		Version   uint8     `json:"version"`
	}

	// APILink stores info about a REST endpoint.
	APILink struct {
		Name  string    `json:"endpoint"`
		Funcs []APIFunc `json:"functions"`
	}

	// APIFunc stores info about a REST method.
	APIFunc struct {
		Name   string `json:"name"`
		Method string `json:"method"`
		URL    string `json:"url"`
		Desc   string `json:"desc"`
	}

	// Annotation stores all information about a user's annotation.
	Annotation struct {
		ID            string    `bson:"id" json:"id"`
		SchemaVersion string    `bson:"annotator_schema_version" json:"annotator_schema_version"`
		Created       time.Time `bson:"created" json:"created"`
		Updated       time.Time `bson:"updated" json:"updated"`
		Text          string    `bson:"text" json:"text"`
		Quote         string    `bson:"quote" json:"quote"`
		URI           string    `bson:"uri" json:"uri"`
		Owner         string    `bson:"user" json:"user"`
		Consumer      string    `bson:"consumer" json:"consumer"`
	}

	// AnnotationRange stores the relative location of an annotation
	AnnotationRange struct {
		Start       string `bson:"start" json:"start"`
		End         string `bson:"end" json:"end"`
		StartOffset uint   `bson:"startOffset" json:"startOffset"`
		EndOffset   uint   `bson:"endOffset" json:"endOffset"`
	}
)

// StoreMetadata stores info about the annotation storage API.
var (
	StoreMetadata = Metadata{
		Message: "Annotation Store API",
		Endpoints: []APILink{
			APILink{
				Name: "Annotation",
				Funcs: []APIFunc{
					APIFunc{
						Name:   "create",
						Method: "PUT",
						URL:    "",
						Desc:   "Create a new annotation",
					},
					APIFunc{
						Name:   "read",
						Method: "GET",
						URL:    "",
						Desc:   "Get an existing annotation",
					},
					APIFunc{
						Name:   "update",
						Method: "PUT",
						URL:    "",
						Desc:   "Update an existing annotation",
					},
					APIFunc{
						Name:   "delete",
						Method: "DELETE",
						URL:    "",
						Desc:   "Delete an annotation",
					},
				},
			},
		},
		Version: 1,
	}

	TestAnnotations = []Annotation{
		Annotation{
			ID:            uuid.NewV4().String(),
			SchemaVersion: "v1.0",
			Created:       time.Now(),
			Updated:       time.Now(),
			Text:          "This is an annotation",
			Quote:         "This is the annotated text",
			URI:           "This is the location",
			Owner:         "alice",
			Consumer:      "backend",
		},
		Annotation{
			ID:            uuid.NewV4().String(),
			SchemaVersion: "v1.0",
			Created:       time.Now(),
			Updated:       time.Now(),
			Text:          "This is an annotation",
			Quote:         "This is the annotated text",
			URI:           "This is the location",
			Owner:         "alice",
			Consumer:      "backend",
		},
	}
)
