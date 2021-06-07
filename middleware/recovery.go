package middleware

import (
	"fmt"
	"ginhello/constant"
	"ginhello/logger"
	"ginhello/model"
	"ginhello/util"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"strconv"
)

/**
 *统一异常处理
 */
func recoverLog(ctx *gin.Context)  {
	defer func() {
		if err := recover();err!=nil{
			errMessage := err
			str :=string(debug.Stack())

			switch val :=err.(type){
			case logger.CustomLog:
				errMessage = util.TrimmedPath(val.File) + ":" + strconv.Itoa(val.Line) + "\t" + val.Message
				str = ""
			}
			message := fmt.Sprintf("%s\t%s\t%s\t%s\r\n%s",
				ctx.Value("requestId").(string), errMessage, ctx.Request.URL.Path, ctx.Request.URL.RawQuery, str)
			//日志打印
			logger.Logger.Error(message)
			//返回统一json
			result := model.NewResult(ctx)
			result.Fail(constant.FailCode,message)
			//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
			ctx.Abort()
			return
		}
	}()
	ctx.Next()
}
