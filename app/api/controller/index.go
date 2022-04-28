package controller

import (
	"fmt"
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/http"
)

type Index struct {
	*Base
}

type GetInput struct {
	Name string `json:"name" form:"name"`
}

func (c Index) Get(ctx *lib.Context) *http.Response {
	inputs := &GetInput{}
	err := ctx.ShouldBind(inputs)
	if err != nil {
		return c.Response(nil, err)
	}
	fmt.Println(inputs.Name)
	return c.Response(nil, nil)
}
