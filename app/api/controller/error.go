package controller

import (
	"gin.go.tpl/kernel/errors"
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Base
}

func (e Error) NoRoute(ctx *gin.Context) {
	response.New().WithIError(errors.Core().RouteNotFound()).Json(ctx)
	return
}
