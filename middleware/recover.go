package middleware

import (
	"gin.go.tpl/lib/code"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	netHttp "net/http"
)

type RecoverMiddleware struct{}

// Broken 恢复内部异常导致的500x错误
func (rm RecoverMiddleware) Broken() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.LogAPI.GetLogger().WithFields(logrus.Fields{
					"panic": r}).Error("panic.log")
				// 打印错误堆栈信息
				// debug.PrintStack()
				c.JSON(netHttp.StatusOK, (&http.Response{}).RespByCode(code.StdErr))
			}
		}()
		c.Next()
	}
}
