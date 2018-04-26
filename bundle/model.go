package bundle

import (
	"bundlehub/common"
	"time"

	"github.com/jinzhu/gorm"
)

type Bundle struct {
	gorm.Model
	Name        string    `json:"name" bson:"name" required:"true"`
	Description string    `json:"description" bson:"description" required:"true"`
	Price1      float32   `json:"price_1" bson:"price_1"`
	Price5      float32   `json:"price_5" bson:"price_5"`
	Price10     float32   `json:"price_10" bson:"price_10"`
	Price25     float32   `json:"price_25" bson:"price_25"`
	Price50     float32   `json:"price_50" bson:"price_50"`
	Start       time.Time `json:"start" bson:"start"`
	Finish      time.Time `json:"finish" bson:"finish"`
	Views       uint      `json:"views" bson:"views"`
	MoneyEarned float32   `json:"money_earned" bson:"money_earned"`
	Purchases   uint      `json:"purchases" bosn:"purchases"`
}

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Omit("created_at", "deleted_at").Save(data).Error
	return err
}
