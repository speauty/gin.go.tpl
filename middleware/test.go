package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type TestMiddleware struct{}

func (td TestMiddleware) Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before action middleware")
		c.Set("name", "middle")
		c.Next()
	}
}

func (td TestMiddleware) After() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		name, _ := c.Get("name")
		status := c.Writer.Status()
		fmt.Println("after action middleware", status, " value:", name)
	}
}
