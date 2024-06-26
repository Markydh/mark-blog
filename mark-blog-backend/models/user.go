package models

import "time"

type User struct {
	Id          int        `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Avatar      string     `json:"avatar"`
	Level       int        `json:"level"`
	CreatedTime *time.Time `json:"created_time"`
	UpdatedTime *time.Time `json:"updated_time"`
}

//用户邮箱注册
type UserEmail struct {
	Id          int        `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	CheckCode   string     `json:"checkCode"`
}
