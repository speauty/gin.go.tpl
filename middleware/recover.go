package middleware

import (
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/log"
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RecoverMiddleware struct{}

// Exec 恢复内部异常导致的500x错误
func (rm RecoverMiddleware) Exec() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.NewLogApi(nil).GetLogger().WithFields(logrus.Fields{
					"panic": r}).Error("panic.log")
				// 打印错误堆栈信息
				//debug.PrintStack()
				(&response.Response{}).WithCode(code.StdErr).Json(ctx)
			}
		}()
		ctx.Next()
	}
}
