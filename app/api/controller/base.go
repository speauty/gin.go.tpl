package controller

import (
	"gin.go.tpl/kernel/errors"
	"gin.go.tpl/kernel/log"
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
)

type Base struct {
}

// Input to receive input from client
func (c *Base) Input(ctx *gin.Context, obj interface{}) errors.Error {
	if err := ctx.ShouldBind(obj); err != nil {
		return errors.BaseError{}.GenFromStdError(err)
	}
	return nil
}

func (c *Base) Response(ctx *gin.Context, resp *response.Response, err errors.Error) *response.Response {
	if err != nil {
		log.NewLogApi(nil).Error(err.Error())
		return nil
	}
	return resp
}
