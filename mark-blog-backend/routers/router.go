/*
routers.go 路由文件,存储api接口
*/
package router

import (
	"mark-blog-backend/global"
	"mark-blog-backend/utils"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	//解决路由跨域问题
	router.Use(utils.Cors())
	gin.SetMode(global.Config.System.Env)
	userRouter(router)
	return router
}
