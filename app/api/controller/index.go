package controller

import (
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/db"
	"gin.go.tpl/lib/http"
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
	ctx.Log.Warn("Hello")
	sqlDb, _ := db.DBAPI.GetDB().DB()
	ctx.Log.Info(sqlDb.Stats())
	if err != nil {
		return c.Response(nil, err)
	}
	return c.Response(nil, nil)
}
