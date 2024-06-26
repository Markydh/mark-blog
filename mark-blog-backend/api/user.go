package api

import (
	"mark-blog-backend/services"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
}

var userService services.UserService

// 获取用户信息
func (UserApi) GetUserInfo(c *gin.Context) {
	userService.GetUserInfo(c)
}

//用户邮箱验证码推送
func (UserApi) SendCodeToUser(c *gin.Context){
	userService.SendCodeToUser(c)
}

// 用户注册
func (UserApi) UserRegister(c *gin.Context) {
	userService.UserRegister(c)
}

func (UserApi) UserLogin(c *gin.Context) {
	userService.UserLogin(c)
}
