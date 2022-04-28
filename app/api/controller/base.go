package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/http"
)

type Base struct {
}

func (c *Base) Input(ctx *lib.Context, obj *interface{}) error {
	err := ctx.ShouldBind(obj)
	if err != nil {
		return err
	}
	return nil
}

func (c *Base) Response(resp *http.Response, err error) *http.Response {
	if err != nil {
		return &http.Response{Code: 400, Msg: err.Error()}
	}
	return resp
}
