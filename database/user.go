package database

import (
	"github.com/tjones879/fake/structs"
	"gopkg.in/mgo.v2"
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
		fn := c.FindId(id).One(&user)
		return fn
	}

	err = withCollection("users", query)
	return
}
