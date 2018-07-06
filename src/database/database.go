package database

import (
	"os"

	"github.com/globalsign/mgo"
)

// Database is a singleton to connect to database
var Database *mgo.Database

func init() {

	mongoUri := os.Getenv("MONGODB_URI")

	session, err := mgo.Dial(mongoUri)
	if err != nil {
		panic(err)
	}

	Database = session.DB(os.Getenv("MONGODB_DATABASE"))

}

// Get a database instance.
func Get() *mgo.Database {
	return Database
}
