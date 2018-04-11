package users

import (
	"github.com/zebresel-com/mongodm"
)

type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`

	UserName     string `json:"username" bson:"username" minLen:"2" maxLen:"30" required:"true"`
	Email        string `json:"email" bson:"email" validation:"email" required:"true"`
	PasswordHash string `json:"-" bson:"passwordHash"`
}
