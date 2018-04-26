package key

import (
	"bundlehub/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func KeyGroup(router *gin.RouterGroup) {
	router.GET("/:id", GetKey)
	router.GET("/", GetKeys)
}

func GetKey(c *gin.Context) {
	id := c.Params.ByName("id")
	var key Key
	if err := common.GetDB().Where("id = ?", id).First(&key).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, key)
	}
}

func GetKeys(c *gin.Context) {
	var keys []Key
	if err := common.GetDB().Find(&keys).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, keys)
	}
}
