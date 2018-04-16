package article

import (
	"bundlehub/common"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ArticleGroup(router *gin.RouterGroup) {
	router.POST("/", RegisterArticle)
}

func RegisterArticle(c *gin.Context) {
	col := common.GetDB().DB("bundlehub").C("article")
	article := Article{}
	if c.BindJSON(&article) == nil {
		log.Println(article.Title)
		log.Println(article.Content)
	}
	err := col.Insert(article)
	if err != nil {
		log.Println(err)
		c.AbortWithError(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusOK, article)
	// userModelValidator := NewUserModelValidator()
	// if err := userModelValidator.Bind(c); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
	// 	return
	// }

	// if err := SaveOne(&userModelValidator.userModel); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
	// 	return
	// }
	// c.Set("my_user_model", userModelValidator.userModel)
	// serializer := UserSerializer{c}
	// c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}
