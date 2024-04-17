/*
routers.go 路由文件,存储api接口
*/
package router

import (
	"mark-blog-backend/controller"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()


	r.GET("/hello",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"name":"mark",
			"age":20,
			"sex":"male",
		})
	})

	//router 配置
	user := r.Group("/user")
	{
		user.GET("/info",controller.GetUserInfo)
		// user.GET("/get", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "Hello Mark")
		// 	ctx.JSON(http.StatusOK,"Hello Mark")
		// })
		// user.POST("/post", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "Hello Post")
		// })
		// user.DELETE("/delete", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "Hello Delete")
		// })
		// user.PUT("/put", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "Hello Put")
		// })

	}
	return r
}
