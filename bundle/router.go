package bundle

import (
	"bundlehub/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func BundleGroup(router *gin.RouterGroup) {
	router.GET("/:id", GetBundle)
	router.GET("/", GetBundles)
}

func GetBundle(c *gin.Context) {
	id := c.Params.ByName("id")
	var bundle Bundle
	if err := common.GetDB().Where("id = ?", id).First(&bundle).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, bundle)
	}
}

func GetBundles(c *gin.Context) {
	var bundles []Bundle
	if err := common.GetDB().Find(&bundles).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, bundles)
	}
}
