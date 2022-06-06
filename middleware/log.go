package middleware

import (
	"gin.go.tpl/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type LogMiddleware struct{}

func (lm LogMiddleware) GoThrough() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawData, _ := c.GetRawData()
		log.LogAPI.GetLogger().WithFields(logrus.Fields{
			"url": c.Request.RequestURI, "method": c.Request.Method, "client_ip": c.ClientIP(), "raw_data": string(rawData),
		}).Info("access.log")
		c.Next()
	}
}
