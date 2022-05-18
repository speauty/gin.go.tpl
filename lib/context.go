package lib

import (
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/http"
	"github.com/gin-gonic/gin"
	"sync"
)

var (
	ContextAPI  *Context
	ContextOnce sync.Once
)

type Context struct {
	*gin.Context
	Config config.Config
}

func NewContextAPI() *Context {
	ContextOnce.Do(func() {
		ContextAPI = &Context{}
		ContextAPI.Config = config.Config{}
	})
	return ContextAPI
}

func (ctx *Context) Init(iniDir string) {
	ctx.Config, _ = ctx.Config.LoadConfig(iniDir)
}

func (ctx *Context) Wrap(handler func(*Context) *http.Response) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		response := handler(&Context{Context: gCtx})
		if response == nil {
			response = response.Default()
		}
		gCtx.JSON(200, response)
	}
}
