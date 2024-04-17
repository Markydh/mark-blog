package router

import (
	"github.com/gin-gonic/gin"
)

func UserAccount(e *gin.Engine) {
	e.POST("/user/register", UserRegister)
	e.POST("/user/login", UserLogin)
	e.DELETE("/user/delete", UserCancle)
	e.POST("/user/changeGrank", ChangeLevel)
	e.POST("/user/updateBlacklistStatus", UpdateBlacklistStatus)
}

//handle

// 用户账号注册
func UserRegister(c *gin.Context) {

}

// 用户账号登陆
func UserLogin(c *gin.Context) {

}

// 用户账号注销
func UserCancle(c *gin.Context) {

}

// 更改黑名单状态
func UpdateBlacklistStatus(c *gin.Context) {

}

// 修改用户登记
func ChangeLevel(c *gin.Context) {

}
