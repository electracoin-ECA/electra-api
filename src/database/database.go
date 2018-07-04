package database

import (
	"os"

	"github.com/globalsign/mgo"
)

// Get a database instance.
func Get() *mgo.Database {

	mongoUri := os.Getenv("MONGODB_URI")

	session, err := mgo.Dial(mongoUri)
	if err != nil {
		panic(err)
	}

	return session.DB(os.Getenv("MONGODB_DATABASE"))
}
