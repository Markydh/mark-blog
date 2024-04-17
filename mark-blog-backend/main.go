package main

import (
	"mark-blog-backend/core"
	"mark-blog-backend/global"
	"mark-blog-backend/router"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()
	global.Log.Warning("warning")
	global.Log.Info("info")
	//连接数据局
	global.DB = core.InitGorm()
	r := router.Router()
	r.Run(":9999")
}
