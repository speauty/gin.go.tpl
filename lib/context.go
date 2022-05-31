package lib

import (
	"gin.go.tpl/lib/cache"
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/lib/log"
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
	Log    *log.Log
	Cache  *cache.Cache
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
	ctx.Log = log.NewLogAPI(ctx.Config.Log)
	ctx.Cache = cache.NewCacheAPI(ctx.Config.Redis)
}

func (ctx *Context) Wrap(handler func(*Context) *http.Response) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		response := handler(&Context{Context: gCtx, Config: ctx.Config, Log: ctx.Log})
		if response == nil {
			response = response.Default()
		}
		gCtx.JSON(200, response)
	}
}
