package router

import "github.com/gin-gonic/gin"

func ArticleManagement(e *gin.Engine) {
	article := e
	article.POST("/article/addArticle", AddArticle)
}

// handle
func AddArticle(c *gin.Context) {

}
