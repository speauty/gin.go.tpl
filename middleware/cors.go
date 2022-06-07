package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CorsMiddleware struct{}

func (cm CorsMiddleware) Exec() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Origin") != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", c.Request.Header.Get("Origin"))
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Header.Get("Access-Control-Request-Headers") != "" {
			c.Writer.Header().Set("Access-Control-Allow-Headers", c.Request.Header.Get("Access-Control-Request-Headers"))
		}

		// to terminate request about option directly
		if c.Request.Method == http.MethodOptions {
			c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST")
			c.Writer.Header().Set("Allow", "OPTIONS, GET, POST")
			c.Writer.Header().Set("Cache-Control", "max-age=604800")
			c.Writer.Header().Set("Content-Length", "0")
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
