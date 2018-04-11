package main

import (
	"bundlehub/common"

	"github.com/gin-gonic/gin"
	"github.com/zebresel-com/mongodm"
)

func initDb(db *mongodm.Connection) {

}

func main() {

	db, _ := common.Init()
	//Migrate(db)
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	r.Run(":1515")
}
