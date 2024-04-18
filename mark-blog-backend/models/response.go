package models

import (
	"github.com/gin-gonic/gin"
)

const (
	Success = 200
	Error   = 400
)

type JsonStruct struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  interface{} `json:"msg"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &JsonStruct{Code: code, Data: data, Msg: msg}
	c.JSON(Success, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}, data interface{}) {
	json := &JsonStruct{Code: code, Msg: msg}
	c.JSON(Error, json)
}
