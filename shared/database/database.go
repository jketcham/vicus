package database

import (
	"log"

	"github.com/jketcham/vicus/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

// MongoDB connection variables returned mgo
var (
	Session *mgo.Session
	Database *mgo.Database
)

// Connect accepts a database url and sets the Session and Database
// connection to the MongoDB server
func Connect(url string) {
	var err error

	Session, err := mgo.Dial(url) // get url
	if err != nil {
		log.Fatalf("Connect: %s\n", err)
	}

	defer Session.Close()

	Session.SetMode(mgo.Monotonic, true)

	Database = Session.DB("vicus_db") // TODO: set in config at startup
}
