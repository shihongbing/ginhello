package main

import (
	"fmt"
	"ginhello/config"
	"ginhello/core"
	"ginhello/logger"
	"ginhello/middleware"
	"ginhello/router"
	"ginhello/util"
	"github.com/gin-gonic/gin"
)

func init() {
	//日志初始化
	//logger.InitLog()
	//logger.SyncLog()
	logger.InitLog(config.LogConfig.Level)
	logger.Logger.Info("系统初始化开始")
	core.ConnectDB()
	logger.Logger.Info("系统初始化结束")
}

func main() {
	logger.Logger.Info("程序启动了!")
	logger.Logger.Info("gopacketTest end!")
	if config.AppConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	//加载中间件
	middleware.LoadMiddleware(engine)
	//加载路由
	router.LoadRoutes(engine)
	//指定运行端口 Run(":8080") 端口前面的冒号不可以省
	engine.Run(config.AppConfig.Port)
	resourceRelease()
	//pwdTest()
}

func pwdTest() {
	passwordOK := "ueba_qwer1234!"
	passwordERR := "adminxx"

	hashStr, err := util.HashAndSalt(passwordOK)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hashStr)

	// 正确密码验证
	check := util.ComparePassword(hashStr, passwordOK)
	if !check {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}

	// 错误密码验证
	check = util.ComparePassword(hashStr, passwordERR)
	if !check {
		fmt.Println("pw wrong")
	} else {
		fmt.Println("pw ok")
	}
}

func resourceRelease() {
	go func() {
		core.DisconnectDB()
	}()
}

//func simpleHttpGet(url string) {
//	logger.Logger.Debugf("Trying to hit GET request for %s", url)
//	resp, err := http.Get(url)
//	if err != nil {
//		logger.Logger.Errorf("Error fetching URL %s : Error = %s", url, err)
//	} else {
//		logger.Logger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
//		resp.Body.Close()
//	}
//}
