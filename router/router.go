package router

import (
	"ginhello/controller"
	"ginhello/middleware"
	"github.com/gin-gonic/gin"
)

func LoadRoutes(router *gin.Engine){
	//authApi := router.Group("/api/auth/")
	//{
	//	authApi.POST("/password", controller.PasswordAuth)
	//}
	//认证api
	authController := controller.AuthController{}
	authController.LoadRouter(router)
	//其他受保护的api
	userController := controller.UserController{}
	userController.LoadRouter(router,middleware.JWTAuth())

}
