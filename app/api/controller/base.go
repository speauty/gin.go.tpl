package controller

import (
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/errors"
	"gin.go.tpl/kernel/log"
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
)

type Base struct {
}

func (controller Base) Input(ctx *gin.Context, obj interface{}) errors.IError {
	if err := ctx.ShouldBind(obj); err != nil {
		return errors.Core().NewFromCode(code.StdInput, err)
	}
	return nil
}

func (controller Base) Response(ctx *gin.Context, resp response.IResponse, err errors.IError) {
	if err != nil {
		log.NewLogApi(nil).Error(err.Error())
		resp.WithIError(err)
	} else {
		resp.WithCode(code.StdOk)
	}
	resp.Json(ctx)
	return
}
