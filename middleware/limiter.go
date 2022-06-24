package middleware

import (
	"gin.go.tpl/kernel/cfg"
	"gin.go.tpl/kernel/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"time"
)

type LimiterMiddleware struct{}

func (lm LimiterMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		l := limiter.NewLimiter(
			rate.Every(time.Duration(cfg.NewCfgApi("").Limiter.GeneratorInterval)*time.Second),
			cfg.NewCfgApi("").Limiter.GeneratorNum, "ALL",
		)
		if !l.Allow() {
			//c.AbortWithStatusJSON(netHttp.StatusOK, (&response.Response{}).RespByCode(code.StdRequestRateExceed))
		} else {
			c.Next()
		}
	}
}
