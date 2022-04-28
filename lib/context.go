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
