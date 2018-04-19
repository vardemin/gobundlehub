package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB field
var DB *gorm.DB

//Init func
func Init() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:pass4thrive@/bundlehub?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	if err != nil {
		fmt.Println("Database connection error: %v", err)
	}

	DB = db
	return DB, err
}

//GetDB retun db instance
func GetDB() *gorm.DB {
	return DB
}
