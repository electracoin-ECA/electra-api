package database

import (
	"log"
	"os"

	"github.com/globalsign/mgo"
)

// Database is a singleton to connect to database
var Database *mgo.Database

func init() {

	mongoUri := os.Getenv("MONGODB_URI")
	if mongoUri == "" {
		log.Println("MONGODB_URI not set")
		os.Exit(5)
	}
	session, err := mgo.Dial(mongoUri)
	if err != nil {
		panic(err)
	}
	db := os.Getenv("MONGODB_DATABASE")
	if db == "" {
		log.Println("MONGODB_DATABASE not set")
		os.Exit(5)
	}
	Database = session.DB(db)

}

// Get a database instance.
func Get() *mgo.Database {
	return Database
}
