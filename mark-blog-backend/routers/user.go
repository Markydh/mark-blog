package router

import (
	"mark-blog-backend/api"

	"github.com/gin-gonic/gin"
)

func userRouter(e *gin.Engine) {
	userApi := api.UserApi{}
	user := e.Group("api/user")
	user.POST("sendCodeToUser",userApi.SendCodeToUser)
	user.POST("userRegister", userApi.UserRegister)
	user.GET("getUserInfo",userApi.GetUserInfo)
	user.GET("userLogin",userApi.UserLogin)
}
