package article

import (
	"bundlehub/common"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleGroup(router *gin.RouterGroup) {
	router.POST("/", RegisterArticle)
	router.GET("/:id", GetArticle)
}

func RegisterArticle(c *gin.Context) {
	article := Article{}
	if c.BindJSON(&article) == nil {
		log.Println(&article)
	}
	err := SaveOne(&article)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, article)
}

func GetArticle(c *gin.Context) {
	id := c.Params.ByName("id")
	var article Article
	if err := common.GetDB().Where("id = ?", id).First(&article).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, article)
	}
}
