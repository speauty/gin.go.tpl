package controller

import (
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	Base
}

func (controller User) Register(ctx *gin.Context) {
	controller.Response(ctx, response.New(), nil)
	return
}

func (controller User) Query(_ *gin.Context) {

}
