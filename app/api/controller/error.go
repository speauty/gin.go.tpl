package controller

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Base
}

func (e Error) NoRoute(ctx *gin.Context) {
	return
}
