package services

import (
	"mark-blog-backend/global"
	"mark-blog-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserService struct {
}

// 获取用户信息
func (UserService) GetUserInfo(c *gin.Context) models.User {
	var user models.User
	err := global.DB.Where("username = ?", c.Query("username")).First(&user).Error
	if err != nil {
		global.Log.Warning("未查询到用户信息")
		return models.User{}
	}
	return models.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Level:    user.Level,
	}
}

// 用户注册
func (UserService) UserRegister(c *gin.Context) {
	var user = models.User{}
	if err := c.Bind(&user); err != nil {
		models.ReturnError(c, 400, "获取json失败", nil)
		return
	}
	var nowTime = time.Now()
	user.CreatedTime = &nowTime
	user.Level = 1
	user.Avatar = "default_avatar.jpg"
	// 将用户对象保存到数据库
	err := global.DB.Create(&user).Error
	if err != nil {
		// 处理错误
		global.Log.Error("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// /用户登陆
func (UserService) UserLogin(c *gin.Context) {
	// 从请求中获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")
	// 查询数据库，检查用户名和密码是否匹配
	var user models.User
	if err := global.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		global.Log.Error("Error finding user:", err)
		models.ReturnError(c, 400, "登陆失败", nil)
	} else {
		models.ReturnSuccess(c, 200, "登陆成功", nil)
	}
}
