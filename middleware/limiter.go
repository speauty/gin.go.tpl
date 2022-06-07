package middleware

import (
	"gin.go.tpl/lib/code"
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/lib/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	netHttp "net/http"
	"time"
)

type LimiterMiddleware struct{}

func (lm LimiterMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := limiter.NewLimiter(
			rate.Every(time.Duration(config.ConfigApi.Limiter.GeneratorInterval)*time.Second),
			config.ConfigApi.Limiter.GeneratorNum, "ALL",
		)
		if !l.Allow() {
			c.AbortWithStatusJSON(netHttp.StatusOK, (&http.Response{}).RespByCode(code.StdRequestRateExceed))
		} else {
			c.Next()
		}
	}
}
