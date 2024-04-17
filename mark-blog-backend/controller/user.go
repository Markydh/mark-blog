package controller

import "github.com/gin-gonic/gin"

func GetUserInfo(c *gin.Context){
	ReturnSuccess(c,200,"Success","Mark",1)
}