package game

import (
	"bundlehub/common"

	"github.com/jinzhu/gorm"
)

type Game struct {
	gorm.Model
	Name         string  `json:"name" bson:"name" required:"true"`
	Description  string  `json:"description" bson:"description" required:"true"`
	Avatar       string  `json:"avatar" bson:"avatar"`
	Developer    string  `json:"developer" bson:"developer"`
	Genre        string  `json:"genre" bson:"genre"`
	Url          string  `json:"url" bson:"url"`
	Price        float32 `json:"price" bson:"price"`
	Achievenemts bool    `json:"achievements" bson:"achievements"`
	Cards        bool    `json:"cards" bson:"cards"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Omit("created_at", "deleted_at").Save(data).Error
	return err
}
