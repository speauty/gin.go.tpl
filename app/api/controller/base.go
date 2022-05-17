package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/code"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/lib/log"
)

type Base struct {
}

// Input to receive input from client, but the struct not matched
// @todo how to receive the struct defined by caller, the type mismatch
func (c *Base) Input(ctx *lib.Context, obj *interface{}) error {
	err := ctx.ShouldBind(obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Base) Response(resp *http.Response, err error) *http.Response {
	if err != nil {
		log.Log{}.Error(err.Error())
		tmpCode := code.StdErr
		if resp != nil && resp.Code != 0 {
			tmpCode = resp.GetCode()
		}
		return resp.RespByCode(tmpCode)
	}
	return resp
}
