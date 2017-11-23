package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	db "github.com/tjones879/fake/database"
	"github.com/tjones879/fake/structs"
	"net/http"
	"strconv"
	"time"
)

var (
	AnnotationPrefix = "/store/annotations"
)

func getUserID(c *gin.Context) (uid string) {
	session := sessions.Default(c)
	userid := session.Get("user-id")
	if userid != nil {
		uid = userid.(string)
	} else {
		uid = ""
	}
	return
}

// RootAnnotate returns metadata.
func RootAnnotate(c *gin.Context) {
	c.JSON(http.StatusOK, structs.StoreMetadata)
}

// IndexAnnotate returns a list of all annotations.
func IndexAnnotate(c *gin.Context) {
	uid := getUserID(c)
	a, err := db.GetAnnotationsByUser(uid)
	fmt.Println("IndexAnnotate:", a)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, a)
}

// CreateAnnotate receives an annotation and saves it.
func CreateAnnotate(c *gin.Context) {
	uid := getUserID(c)
	session := sessions.Default(c)
	var annotate structs.Annotation
	err := c.ShouldBindJSON(&annotate)
	if err == nil && session.Get("user-id") != nil {
		annotate.Created = time.Now()
		annotate.Updated = time.Now()
		annotate.ID = uuid.NewV4().String()
		annotate.Owner = uid
		annotate.SchemaVersion = "v1.0"
		saved, err := json.Marshal(annotate)
		if err != nil {
			fmt.Println("CreateAnnotate:", err)
		} else {
			fmt.Println("Inserting: " + string(saved[:]))
			db.InsertAnnotation(annotate)
		}
	} else {
		fmt.Println("CreateAnnotate:(err)", err, "(userid)", session.Get("user-id"))
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

// UpdateAnnotate updates an existing annotation with new time and text.
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

// DeleteAnnotate removes an existing annotation.
func DeleteAnnotate(c *gin.Context) {
	id := c.Param("id")

	db.DeleteAnnotation(id)

	c.Status(http.StatusNoContent)
}

// SearchAnnotate returns a set of annotations.
func SearchAnnotate(c *gin.Context) {
	uid := getUserID(c)
	uri := c.Query("uri")
	limit := c.DefaultQuery("limit", "20")
	lim, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err)
		lim = 20
	}
	offset := c.DefaultQuery("offset", "0")
	skip, err := strconv.Atoi(offset)
	if err != nil {
		fmt.Println(err)
		skip = 0
	}

	a, num, err := db.GetAnnotationByURI(uid, uri, lim, skip)
	if err != nil {
		fmt.Println(err)
	}
	response := structs.AnnotationSearch{
		Total:   uint(num),
		Results: a,
	}
	c.JSON(http.StatusOK, response)
}

// PageURI returns the correct page URI given an id.
func PageURI(c *gin.Context) {
	fileID := c.Query("id")
	page := db.GetSavedPageByFile(fileID)
	c.JSON(http.StatusOK, struct {
		URL string `json:"url"`
	}{page.Location})
}
