package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code  int         `json:"code"`
	Msg   interface{} `json:"msg"`
	Data  interface{} `json:"data"`
	Count int         `json:"count"`
}

func ReturnSuccess(c *gin.Context, code int, msg interface{}, data interface{}, count int) {
	json := &JsonStruct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(http.StatusOK, json)
}

func ReturnError(c *gin.Context, code int, msg interface{}) {
	json := &JsonStruct{Code: code, Msg: msg}
	c.JSON(200, json)
}
