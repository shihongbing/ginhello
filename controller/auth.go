package controller

import (
	"ginhello/constant"
	"ginhello/logger"
	"ginhello/middleware"
	"ginhello/model"
	"ginhello/util"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)
type AuthController struct {
	BaseController
}
//注册该controller的路由方法
func (controller * AuthController) LoadRouter(engine *gin.Engine,middlewares ...gin.HandlerFunc){
	route := engine.Group("/api/auth/")
	for _,m:= range middlewares{
		route.Use(m)
	}
	route.POST("/pwd", controller.PasswordAuth)
}

//帐号密码认证
func (controller *AuthController)PasswordAuth(ctx * gin.Context)  {
	var result = model.NewResult(ctx)
	var authParam model.AuthParam
	//解析参数
	err:= ctx.BindJSON(&authParam)
	if err == nil{
		var user model.User
		rows,_ := userService.FindByAccount(ctx, authParam.Username,&user)
		if rows <= 0 {
			result.Fail(constant.FailCode, "帐号名或密码错误")
			return
		}
		if util.ComparePassword(user.Password,authParam.Password){
			generateToken(ctx,user,result)
		}else{
			result.Fail(constant.FailCode, "帐号名或密码错误")
		}
	}else{
		result.Fail(constant.FailCode, "认证参数解析错误:"+err.Error())
	}
}


// token生成器
func generateToken(c *gin.Context, user model.User,result *model.Result) {
	// 构造SignKey: 签名和解签名需要使用一个值
	j := middleware.NewJWT()

	// 构造用户claims信息(负荷)
	claims := middleware.CustomClaims{
		Id:      user.Id,
		Account: user.Account,
		Name:    user.UserName,
		Email:   user.Email,
		StandardClaims: jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 签名过期时间
			Issuer:    "bgbiao.top",                    // 签名颁发者
		},
	}
	// 根据claims生成token对象
	token, err := j.CreateToken(claims)

	if err != nil {
		result.Fail(constant.FailCode,err.Error())
	}
	logger.Logger.Info("generateToken：",token)
	result.Success(constant.SuccessCode,"认证成功",token)
}