package router

import (
	"gin.go.tpl/app/api/controller"
	"gin.go.tpl/lib"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r Router) GetRouters(ctx *lib.Context, engine *gin.Engine) *gin.Engine {
	router := engine
	router.GET("/ping", ctx.Wrap(controller.Index{}.Get))
	return router
}
