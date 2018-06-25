package database

import (
	"os"

	"github.com/globalsign/mgo"
)

// Get a database instance.
func Get() *mgo.Database {
	session, err := mgo.Dial(os.Getenv("MONGODB_URI"))
	if err != nil {
		panic(err)
	}

	return session.DB(os.Getenv("MONGODB_DATABASE"))
}
