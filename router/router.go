package router

import (
	"gin.go.tpl/app/api/controller"
	"github.com/gin-gonic/gin"
)

type Router struct{}

func (r Router) GetRouters(engine *gin.Engine) *gin.Engine {
	router := engine
	// 定义404处理句柄
	router.NoRoute(controller.Error{}.NoRoute)

	router.GET("/register", controller.User{}.Register)
	router.GET("/query", controller.User{}.Query)

	return router
}
