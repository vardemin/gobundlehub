package game

import (
	"bundlehub/common"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GameGroup(router *gin.RouterGroup) {
	router.GET("/:id", GetGame)
	router.GET("/", GetGames)
}

func GetGame(c *gin.Context) {
	id := c.Params.ByName("id")
	var game Game
	if err := common.GetDB().Where("id = ?", id).First(&game).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, game)
	}
}

func GetGames(c *gin.Context) {
	var games []Game
	if err := common.GetDB().Find(&games).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, games)
	}
}
