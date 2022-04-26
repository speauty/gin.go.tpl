package lib

import (
	"gin.go.tpl/lib/http"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func (ctx Context) Wrap(handler func(*Context) *http.Response) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		response := handler(&Context{Context: gCtx})

		gCtx.JSON(200, response)

	}
}

// func convert(f func(*context.Context) *response.Response) gin.HandlerFunc {
//
//	return func(c *gin.Context) {
//
//		resp := f(&context.Context{Context: c})
//
//		data := resp.GetData()
//
//		switch item := data.(type) {
//
//		case string:
//
//			c.String(200, item)
//
//		case gin.H:
//
//			c.JSON(200, item)
//
//		}
//
//	}
//
//}
