/*
routers.go 路由文件,存储api接口
*/
package router

import (
	"mark-blog-backend/global"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(global.Config.System.Env)
	userRouter(router)
	return router
}
