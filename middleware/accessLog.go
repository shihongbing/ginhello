package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

/**
 *访问日志记录
 */
func accessLog(logger *zap.SugaredLogger) gin.HandlerFunc{
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Request.ParseForm()
		requestId, _ := c.Value("requestId").(string)
		logger.Info("request ",
			zap.String("requestId", requestId),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()),
			zap.String("path", path),
			zap.String("query", query),
			zap.Any("post", c.Request.PostForm),
		)
		c.Next()
		end := time.Now()
		cost:= end.Sub(start)
		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				logger.Error(e)
			}
		} else {
			response, _ := c.Get("response")

			logger.Infof("requestId %s response :http_status= %d data= %v cost=%d",
				requestId,c.Writer.Status(),response,cost)
		}
	}
}
