package router

import (
	"gin.go.tpl/app/api/controller"
	"gin.go.tpl/lib"
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (r Router) GetRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", lib.Context{}.Wrap(controller.Index{}.Get))
	return router
}
