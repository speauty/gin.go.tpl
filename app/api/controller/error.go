package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/errors"
	"gin.go.tpl/lib/http"
)

type Error struct {
	Base
}

func (e Error) NoRoute(ctx *lib.Context) *http.Response {
	return e.Response(ctx, nil, errors.SysError{}.RouteNotFound())
}
