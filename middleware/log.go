package middleware

import (
	"gin.go.tpl/kernel/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct{}

func (lm LogMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 调整成异步处理
		go func() {
			log.NewLogApi(nil).GetLogger().WithFields(logrus.Fields{
				"url": c.Request.RequestURI, "method": c.Request.Method,
				"client": c.ClientIP(), "user-agent": c.Request.UserAgent(),
				"referer": c.Request.Referer(), "host": c.Request.Host,
			}).Info("access.log")
		}()
		c.Next()
	}
}
