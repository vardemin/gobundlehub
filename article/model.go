package article

import (
	"bundlehub/common"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `json:"title" bson:"title" required:"true"`
	Content string `json:"content" bson:"content" required:"true"`
	Image   string `json:"image" bson:"image"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Omit("created_at", "deleted_at").Save(data).Error
	return err
}
