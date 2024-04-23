package services

import (
	"errors"
	"fmt"
	"mark-blog-backend/global"
	"mark-blog-backend/models"
	"mark-blog-backend/utils"
	"net/http"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
}

// 获取用户信息
func (UserService) GetUserInfo(c *gin.Context) {
	var user models.User
	err := global.DB.Where("username = ?", c.Query("username")).First(&user).Error
	if err != nil {
		models.ReturnError(c, 400, "未查询到用户信息", nil)
		global.Log.Warning("未查询到用户信息")
		return
	}
	var userInfo models.User = models.User{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Avatar:   user.Avatar,
		Level:    user.Level,
	}
	models.ReturnSuccess(c, 200, "查询成功", userInfo)
}

// @Title GetValidateCode
// @Description  发送邮箱验证码 并存入redis（5分钟有效时间）
// @Author Markydh 2024-04-21 20:14:20
// @Param c type description
func (UserService) SendCodeToUser(c *gin.Context) {

	var user models.User
	if err := c.Bind(&user); err != nil {
		models.ReturnError(c, 400, "获取json失败", nil)
		return
	}

	// 检查邮箱是否已经存在于数据库
	var existingUser models.User
	if err := global.DB.First(&existingUser, "email = ?", user.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			global.Log.Infof("邮箱未注册")
		}
	} else {
		models.ReturnError(c, 400, "邮箱已注册", nil)
		fmt.Println("邮箱已注册")
		return
	}

	//将验证码发到用户指定邮箱 并返回验证码
	vCode, err := utils.SendCodeValidate(user)
	if err != nil {
		global.Log.Warning(err)
		models.ReturnError(c, 400, "验证码发送失败", nil)
		return
	}

	// 验证码存入redis 并设置过期时间5分钟
	// 从连接池获取连接
	conn := global.RedisPool.Get()
	//关闭连接
	defer conn.Close()
	replySet, err := conn.Do("hset", "user:code", user.Username, vCode)
	if err != nil {
		panic(err)
	}
	conn.Do("expire", "user:code", 300)

	global.Log.Info(replySet)
	models.ReturnSuccess(c, 200, "验证码已发送到目标邮箱", nil)
}

// @Title GetValidateCode
// @Description  读取用户注册信息，判断验证码是否正确，正确则创建用户账号
// @Author Markydh 2024-04-22 16:40:32
// @Param c type description
func (UserService) UserRegister(c *gin.Context) {
	var userEmail = models.UserEmail{}
	var user = models.User{}
	if err := c.Bind(&userEmail); err != nil {
		models.ReturnError(c, 400, "获取json失败", nil)
		return
	}

	// 从redis中提取验证码，判断用户提交的验证码是否正确

	conn := global.RedisPool.Get()
	defer conn.Close()
	vCode, err := redis.String(conn.Do("hget", "user:code", userEmail.Username))
	if err != nil {
		panic(err)
	}
	if vCode == "" {
		models.ReturnError(c, 400, "未收到验证码", nil)
		return
	}
	if vCode != userEmail.CheckCode {
		models.ReturnError(c, 400, "验证码错误", nil)
		return
	}
	//如果验证码正确，则注册用户账号

	//hash加密
	hashPwd, err := utils.HashPassword(userEmail.Password)
	if err != nil {
		global.Log.Warning("加密失败")
	}
	user.Username = userEmail.Username
	user.Email = userEmail.Email
	user.Password = hashPwd
	var nowTime = time.Now()
	user.CreatedTime = &nowTime
	user.Level = 1
	user.Avatar = "default_avatar.jpg"

	// 将用户对象保存到数据库
	if err := global.DB.Create(&user).Error; err != nil {
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
	if err := global.DB.Where("username = ?", username).First(&user).Error; err != nil {
		global.Log.Error("Error finding user:", err)
		models.ReturnError(c, 400, "登陆失败，用户名不存在", nil)
	} else {
		if utils.ComparePasswords(user.Password, password) {
			models.ReturnSuccess(c, 200, "登陆成功", nil)
		} else {
			models.ReturnSuccess(c, 400, "登陆失败，密码不匹配", nil)
		}
	}
}
