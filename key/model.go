package key

import (
	"bundlehub/common"

	"github.com/jinzhu/gorm"
)

type Key struct {
	gorm.Model
	GameID    int    `json:"game" bson:"game" gorm:"index"`
	Key       string `json:"key" bson:"key"`
	Purchased bool   `json:"purchased" bson:"purchased"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Omit("created_at", "deleted_at").Save(data).Error
	return err
}
