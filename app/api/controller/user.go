package controller

import (
	"gin.go.tpl/db/dao"
	"gin.go.tpl/lib"
	"gin.go.tpl/lib/code"
	"gin.go.tpl/lib/http"
	"gin.go.tpl/service"
)

type User struct {
	Base
}

type userRegisterInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Passwd   string `json:"passwd" binding:"required"`
}

func (u User) Register(ctx *lib.Context) *http.Response {
	inputs := &userRegisterInput{}
	if err := u.Input(ctx, inputs); err != nil {
		return u.Response(ctx, (&http.Response{}).RespByCode(code.StdInput), err)
	}
	if err := (service.UserService{}).Register(ctx, &dao.UserDao{Nickname: inputs.Nickname, Passwd: inputs.Passwd}); err != nil {
		return u.Response(ctx, nil, err)
	}
	return u.Response(ctx, nil, nil)
}

type userQueryInput struct {
	Id int64 `form:"id" binding:"required"`
}

func (u User) Query(ctx *lib.Context) *http.Response {
	inputs := &userQueryInput{}
	if err := u.Input(ctx, inputs); err != nil {
		return u.Response(ctx, (&http.Response{}).RespByCode(code.StdInput), err)
	}
	userDao := &dao.UserDao{Id: inputs.Id}
	if err := (service.UserService{}).Query(ctx, userDao); err != nil {
		return u.Response(ctx, nil, err)
	}
	return u.Response(ctx, (&http.Response{Data: userDao}).RespByCode(code.StdOk), nil)
}
