package main

import (
	"bundlehub/article"
	"bundlehub/bundle"
	"bundlehub/common"
	"bundlehub/game"
	"bundlehub/key"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&article.Article{})
	db.AutoMigrate(&game.Game{})
	db.AutoMigrate(&bundle.Bundle{})
	db.AutoMigrate(&key.Key{})
}

func main() {
	db, _ := common.Init()
	Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	article.ArticleGroup(v1.Group("/article"))
	game.GameGroup(v1.Group("/game"))
	bundle.BundleGroup(v1.Group("/bundle"))
	key.KeyGroup(v1.Group("/key"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":1515")
}
