package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/errors"
	"gin.go.tpl/lib/http"
)

type Base struct {
}

// Input to receive input from client
func (c *Base) Input(ctx *lib.Context, obj interface{}) errors.Error {
	if err := ctx.ShouldBind(obj); err != nil {
		return errors.BaseError{}.GenFromStdError(err)
	}
	return nil
}

func (c *Base) Response(ctx *lib.Context, resp *http.Response, err errors.Error) *http.Response {
	if err != nil {
		ctx.Log.Error(err.Error())
		tmpCode := err.GetCode()
		tmpMsg := err.Error()
		if resp != nil && resp.Code != 0 {
			tmpCode = resp.GetCode()
			tmpMsg = resp.GetCode().GetMsg()
		}
		return resp.RespByCodeAndMsg(tmpCode, tmpMsg)
	}
	return resp
}
