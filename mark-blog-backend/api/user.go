package api

import (
	"mark-blog-backend/models"
	"mark-blog-backend/services"

	"github.com/gin-gonic/gin"
)



type UserApi struct {
}

var userService services.UserService



//获取用户信息
func (UserApi) GetUserInfo(c *gin.Context){
	data:=userService.GetUserInfo(c)
	models.ReturnSuccess(c,200,"查询成功",data)
}


//用户注册
func (UserApi) UserRegister(c *gin.Context) {
	userService.UserRegister(c)
}


func (UserApi)UserLogin(c *gin.Context){
	 userService.UserLogin(c)
}
