package utils

import (
	"mark-blog-backend/global"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 使用 Bcrypt 算法生成密码哈希值
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
		global.Log.Warning("加密失败")
        return "", err
    }
    return string(hashedPassword), nil
}


// ComparePasswords 比较输入的密码与哈希值是否匹配
func ComparePasswords(hashedPassword, password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
    return err == nil
}