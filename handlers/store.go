package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	db "github.com/tjones879/fake/database"
	"github.com/tjones879/fake/structs"
	"net/http"
	"time"
)

var (
	AnnotationPrefix = "/store/annotations"
)

/*
	// Annotation stores all information about a user's annotation.
	Annotation struct {
		ID            string    `json:"id"`
		SchemaVersion string    `json:"annotator_schema_version"`
		Created       time.Time `json:"created"`
		Updated       time.Time `json:"updated"`
		Text          string    `json:"text"`
		Quote         string    `json:"quote"`
		URI           string    `json:"uri"`
		Owner         string    `json:"user"`
		Consumer      string    `json:"consumer"`
	}

	// AnnotationRange stores the relative location of an annotation
	AnnotationRange struct {
		Start       string `json:"start"`
		End         string `json:"end"`
		StartOffset uint   `json:"startOffset"`
		EndOffset   uint   `json:"endOffset"`
	}
*/

// RootAnnotate returns metadata.
func RootAnnotate(c *gin.Context) {
	c.JSON(http.StatusOK, structs.StoreMetadata)
}

// IndexAnnotate returns a list of all annotations.
func IndexAnnotate(c *gin.Context) {
	c.JSON(http.StatusOK, structs.TestAnnotations)
}

// CreateAnnotate receives an annotation and saves it.
func CreateAnnotate(c *gin.Context) {
	var annotate structs.Annotation
	if c.ShouldBindJSON(&annotate) == nil {
		annotate.Created = time.Now()
		annotate.Updated = time.Now()
		annotate.ID = uuid.NewV4().String()
		saved, _ := json.Marshal(annotate)
		fmt.Println("Inserting: " + string(saved[:]))
		db.InsertAnnotation(annotate)
	}
	c.Redirect(http.StatusSeeOther, AnnotationPrefix+"/"+annotate.ID)
}

// ReadAnnotate returns the annotation with the given id.
func ReadAnnotate(c *gin.Context) {
	id := c.Param("id")

	a, err := db.GetAnnotationByID(id)
	if err != nil {
		fmt.Println("ReadAnnotate:", err)
	}
	c.JSON(http.StatusOK, a)
}

func UpdateAnnotate(c *gin.Context) {
	var annotate structs.Annotation
	if c.ShouldBindJSON(&annotate) == nil {
		annotate.Updated = time.Now()
		saved, _ := json.Marshal(annotate)
		fmt.Println("Inserting: " + string(saved[:]))
		db.UpdateAnnotation(annotate)
	}

	c.Redirect(http.StatusSeeOther, AnnotationPrefix+"/"+annotate.ID)
}

func DeleteAnnotate(c *gin.Context) {
	id := c.Param("id")

	db.DeleteAnnotation(id)

	c.Status(http.StatusNoContent)
}
