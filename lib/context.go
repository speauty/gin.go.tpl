package lib

import (
	"gin.go.tpl/lib/cache"
	"gin.go.tpl/lib/config"
	"gin.go.tpl/lib/db"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/lib/log"
	"github.com/gin-gonic/gin"
	netHttp "net/http"
	"sync"
)

var (
	ContextApi  *Context
	ContextOnce sync.Once
)

type Context struct {
	*gin.Context
	Config *config.Config
	Log    *log.Log
	DB     *db.DB
	Cache  *cache.Cache
}

func NewContextApi() *Context {
	ContextOnce.Do(func() {
		ContextApi = &Context{}
	})
	return ContextApi
}

func (ctx *Context) Init() {
	ctx.Config = config.ConfigApi
	ctx.Log = log.LogApi
	ctx.DB = db.DBApi
	ctx.Cache = cache.CacheApi
}

func (ctx *Context) Wrap(handler func(*Context) *http.Response) gin.HandlerFunc {
	return func(gCtx *gin.Context) {
		response := handler(&Context{Context: gCtx, Config: ctx.Config, Log: ctx.Log, DB: ctx.DB})
		if response == nil {
			response = response.Default()
		}
		gCtx.JSON(netHttp.StatusOK, response)
	}
}
