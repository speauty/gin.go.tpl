package router

import (
	"gin.go.tpl/app/api/controller"
	"gin.go.tpl/lib"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r Router) GetRouters(ctx *lib.Context, engine *gin.Engine) *gin.Engine {
	router := engine
	// 定义404处理句柄
	router.NoRoute(ctx.Wrap(controller.Error{}.NoRoute))

	router.POST("/register", ctx.Wrap(controller.User{}.Register))
	router.GET("/query", ctx.Wrap(controller.User{}.Query))

	return router
}
