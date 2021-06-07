package model

import (
	"ginhello/constant"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Ctx *gin.Context
}

type ResultJson struct {
	Code int `json:"code"` //提示代码
	Msg string `json:"message"` //提示信息
	Data interface{} `json:"data"` //数据
}

func NewResult(ctx *gin.Context)  *Result{
	return & Result{Ctx: ctx}
}

//默认成功返回json
func (result * Result)DefaultSuccess(data interface{})  {
	if data == nil {
		data = gin.H{}
	}
	json := ResultJson{Code: constant.SuccessCode,Msg: constant.SuccessMessage,Data: data}
	result.Ctx.Set("response", json)
	result.Ctx.JSON(http.StatusOK,json)
}
//自定义成功返回
func  (result * Result) Success(code int,msg string,data interface{})  {
	json := ResultJson{Code: code,Msg: msg,Data: data}
	//gin.Context中设置response属性,accessLog中间介中取值日志记录
	result.Ctx.Set("response", json)
	result.Ctx.JSON(http.StatusOK,json)
}
//失败返回json
func  (result * Result) Fail(code int,msg string)  {
	json := ResultJson{Code: code,Msg: msg,Data: gin.H{}}
	result.Ctx.Set("response", json)
	result.Ctx.JSON(http.StatusOK,json)
}
//默认失败返回
func  (result * Result) DefaultFail()  {
	json := ResultJson{Code: constant.FailCode,Msg: constant.FailMessage,Data: gin.H{}}
	result.Ctx.Set("response", json)
	result.Ctx.JSON(http.StatusOK,json)
}



