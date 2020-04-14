package helpers

import (
	"log"

	"gopkg.in/mgo.v2"
)

var db *mgo.Database

func InitDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("todos_api")
}

func Collection(index int) *mgo.Collection {
	if index == 0 {
		return db.C("users")
	}
	if index == 1 {
		return db.C("todos")
	}
	return db.C("dump")
}
