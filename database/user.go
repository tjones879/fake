package database

import (
	"fmt"
	"github.com/tjones879/fake/structs"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ensureUsersIndex(s *mgo.Session) error {
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

	return withCollection("users", query)
}

// InsertUser attempts to insert a new user into the users collection
func InsertUser(q structs.User) (insertError error) {
	query := func(c *mgo.Collection) error {
		fn := c.Insert(q)
		return fn
	}

	return withCollection("users", query)
}

// GetUserByID returns a user with the given id.
func GetUserByID(id string) (user structs.User, err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(bson.M{"id": id}).One(&user)
		return fn
	}

	err = withCollection("users", query)
	return
}

// GetUserPages TODO
func GetUserPages(u structs.User) (files []structs.FileStorage) {
	for _, p := range u.FileIDs {
		f, err := GetFileByID(p)
		fmt.Println("fileID", p)
		if err != nil {
			fmt.Println("GetUserPages:", err)
		} else {
			fmt.Println(f)
			files = append(files, f)
		}
	}
	return
}

// AddFileToUser TODO
func AddFileToUser(uid string, f structs.FileStorage) error {
	// pages
	match := bson.M{"id": uid}
	// TODO find a better way to select `files` field from doc
	change := bson.M{"$push": bson.M{"files": f.UID}}
	return usingCollection("users").Update(match, change)
}
