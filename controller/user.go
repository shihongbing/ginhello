package controller

import (
	"ginhello/constant"
	"ginhello/model"
	"ginhello/service"
	"github.com/gin-gonic/gin"
)

var userService = service.UserService{}

type UserController struct {
	BaseController
}

//注册该controller的路由方法
func (controller * UserController) LoadRouter(engine *gin.Engine,middlewares ...gin.HandlerFunc){
	route := engine.Group("/api/user/")
	for _,m:= range middlewares{
		route.Use(m)
	}
	route.GET("/get/:id", controller.Get)
}

//根据ID查询用户
func (controller *UserController) Get(ctx * gin.Context)  {
	var user model.User
	var result = model.NewResult(ctx)
	id := ctx.Param("id")
	if id==""{
		result.Fail(constant.FailCode,"参数解析错误")
		return
	}
	rows,_ := userService.Get(ctx,id,&user)
	if rows <= 0 {
		result.Fail(constant.FailCode, "用户不存在")
		return
	}
	result.DefaultSuccess(user)
}