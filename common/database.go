package common

import (
	"fmt"

	"github.com/zebresel-com/mongodm"
)

//DB field
var Connection *mongodm.Connection

//Init func
func Init() (*mongodm.Connection, error) {
	dbConfig := &mongodm.Config{
		DatabaseHosts: []string{"127.0.0.1"},
		DatabaseName:  "bundlehub",
	}

	connection, err := mongodm.Connect(dbConfig)

	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}
	return connection, err
}

//GetDB retun db instance
func GetDB() *mongodm.Connection {
	return Connection
}
