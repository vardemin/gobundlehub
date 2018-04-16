package main

import (
	"bundlehub/article"
	"bundlehub/common"

	"github.com/gin-gonic/gin"
)

func main() {
	db, _ := common.Init()
	//db, _ := common.Init()
	//Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")
	article.ArticleGroup(v1.Group("/article"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":1515")
}
