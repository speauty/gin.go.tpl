package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/service"
)

type Index struct {
	*Base
}

type GetInput struct {
	Name string `json:"name" form:"name" binding:"required"`
}

func (c Index) Get(ctx *lib.Context) *http.Response {
	inputs := &GetInput{}
	err := ctx.ShouldBind(inputs)
	service.MigratorService{}.SyncTables(ctx)
	if err != nil {
		return c.Response(nil, err)
	}
	return c.Response(nil, nil)
}
