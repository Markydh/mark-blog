package main

import (
	"mark-blog-backend/core"
	"mark-blog-backend/global"
	router "mark-blog-backend/routers"
)

func main() {
	//读取配置文件
	core.InitConfig()
	//初始化日志
	global.Log = core.InitLogger()
	//连接数据局
	global.DB = core.InitGorm()
	//连接redis
	// global.Redis=core.InitRedis()
	global.RedisPool = core.InitRedis()
	//配置路由
	r := router.SetUpRouter()
	addr := global.Config.System.Addr()
	// global.Log.Infof("mark-blog-backend运行在:%s", addr)
	r.Run(addr)
}
