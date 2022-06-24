package controller

import (
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	Base
}

type userRegisterInput struct {
	Nickname string `json:"nickname" binding:"required"`
	Passwd   string `json:"passwd" binding:"required"`
}

func (u User) Register(ctx *gin.Context) {
	response.Resp().SetCode(code.StdRequestRateExceed).(response.IResponse).SetMsg("测试一下").(response.IResponse).Json(ctx)
	//resp.Resp(ctx)
	return
}

type userQueryInput struct {
	Id int64 `form:"id" binding:"required"`
}

func (u User) Query(ctx *gin.Context) {

}
