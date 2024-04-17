package main

import (
	"mark-blog-backend/dao"
	"mark-blog-backend/router"
)

func main() {
	dao.Init()
	dao.Query()
	r := router.Router()
	r.Run(":9999")
}
