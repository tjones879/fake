package database

import (
	"github.com/tjones879/fake/structs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var annotations = "annotations"

func ensureAnnotationsIndex(s *mgo.Session) error {
	session := s.Copy()
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	query := func(c *mgo.Collection) error {
		return c.EnsureIndex(index)
	}

	return withCollection(annotations, query)
}

func ensureAnnotationUserIndex(s *mgo.Session) error {
	session := s.Copy()
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"owner"},
		Unique:     false,
		DropDups:   false,
		Background: true,
		Sparse:     true,
	}

	query := func(c *mgo.Collection) error {
		return c.EnsureIndex(index)
	}

	return withCollection(annotations, query)
}

// GetAnnotationByID returns an annotation with the given id.
func GetAnnotationByID(id string) (a structs.Annotation, err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"id": id}).One(&a)
		return fn
	}

	err = withCollection(annotations, query)
	return
}

// GetAnnotationsByUser returns all annotations that a specific user has generated.
func GetAnnotationsByUser(user string) (a []structs.Annotation, err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"owner": user}).All(&a)
		return fn
	}

	err = withCollection(annotations, query)
	return
}

// GetAnnotationByURI returns some annotations for the specified uri.
func GetAnnotationByURI(uid, uri string, limit, skip int) (a []structs.Annotation, num int, err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"uri": uri, "user": uid}).Skip(skip).Limit(limit).All(&a)
		return fn
	}

	_ = withCollection(annotations, query)
	num, err = usingCollection(annotations).Find(bson.M{"uri": uri, "user": uid}).Count()
	return
}

// InsertAnnotation inserts a new annotation
func InsertAnnotation(a structs.Annotation) (insertError error) {
	query := func(c *mgo.Collection) error {
		fn := c.Insert(a)
		return fn
	}

	return withCollection(annotations, query)
}

// UpdateAnnotation updates an existing annotation.
func UpdateAnnotation(a structs.Annotation) (insertError error) {
	query := func(c *mgo.Collection) error {
		fn := c.Update(bson.M{"id": a.ID}, a)
		return fn
	}

	return withCollection(annotations, query)
}

// DeleteAnnotation deletes an existing annotation.
func DeleteAnnotation(id string) error {
	query := func(c *mgo.Collection) error {
		fn := c.Remove(bson.M{"id": id})
		return fn
	}

	return withCollection(annotations, query)
}
