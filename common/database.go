package common

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

//DB field
var Session *mgo.Session

//Init func
func Init() (*mgo.Session, error) {

	Session, err := mgo.Dial("mongodb://127.0.0.1:27017")

	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}

	return Session, err
}

//GetDB retun db instance
func GetDB() *mgo.Session {
	return Session
}
