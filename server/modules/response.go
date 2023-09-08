package modules

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response resp struct
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var (
	SUCCESS    = 0
	ERROR      = 1
	ERRORLOGIN = 2
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OKResp(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OKWithMsgResp(msg string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, msg, c)
}

func OKWithDataResp(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OKWithAllResp(data interface{}, msg string, c *gin.Context) {
	Result(SUCCESS, data, msg, c)
}

func ErrResp(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "系统错误，请联系管理员", c)
}

func ErrWithMsg(msg string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, msg, c)
}

func OKNothing(c *gin.Context) {
	c.JSON(SUCCESS, Response{Code: 0, Data: nil, Msg: "success"})
}

func ErrOtWithMsg(msg string, c *gin.Context) {
	Result(ERRORLOGIN, map[string]interface{}{}, msg, c)
}
