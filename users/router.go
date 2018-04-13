package users

import (
	"bundlehub/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveOne(data interface{}) error {
	db := common.GetDB()
	err := db.Save(data).Error
	return err
}

func Register(c *gin.Context) {
	if err := c.Bind(,) gin.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "GTFO")
		return
	}

	if err := SaveOne(&userModelValidator.userModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModelValidator.userModel)
	serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
