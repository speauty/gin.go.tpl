package router

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/http"
	"github.com/gin-gonic/gin"
)

//router := gin.Default()
//router.GET("/ping", func(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"message": "pong",
//	})
//})

type Router struct {
}

func (r Router) GetRouters() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", lib.Context{}.Wrap(func(c *lib.Context) *http.Response {
		resp := &http.Response{Code: 200, Msg: "方式"}
		return resp
	}))
	return router
}
