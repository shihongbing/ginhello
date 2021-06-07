package middleware

import (
	"ginhello/logger"
	"github.com/gin-gonic/gin"
)

func LoadMiddleware(engine *gin.Engine) ()  {
	engine.Use(generateRequestId)
	engine.Use(accessLog(logger.Logger))

	engine.Use(recoverLog)
}
