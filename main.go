package main

import (
	"bundlehub/article"
	"bundlehub/common"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&article.Article{})
}

func main() {
	db, _ := common.Init()
	Migrate(db)
	defer db.Close()

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
