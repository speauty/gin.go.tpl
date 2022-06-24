package response

import (
	"gin.go.tpl/kernel/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IResponse interface {
	GetCode() code.Code
	GetMsg() string
	GetData() interface{}
	SetCode(argCode code.Code) interface{}
	SetMsg(msg string) interface{}
	SetData(data interface{}) interface{}
	Default() interface{}
	WithCode(argCode code.Code) interface{}
	WithCodeAndMsg(argCode code.Code, msg string) interface{}
	WithError(argCode code.Code, err error) interface{}
	Json(ctx *gin.Context)
}

func Resp() *Response {
	return &Response{}
}

type Response struct {
	Code code.Code   `json:"c"`
	Msg  string      `json:"m"`
	Data interface{} `json:"d"`
}

func (r Response) GetCode() code.Code {
	return r.Code
}

func (r Response) GetMsg() string {
	return r.Msg
}

func (r Response) GetData() interface{} {
	return r.Data
}

func (r *Response) SetCode(argCode code.Code) interface{} {
	r.Code = argCode
	return r
}

func (r *Response) SetMsg(msg string) interface{} {
	r.Msg = msg
	return r
}

func (r *Response) SetData(data interface{}) interface{} {
	r.Data = data
	return r
}

func (r *Response) Default() interface{} {
	r.SetCode(code.StdOk)
	return r
}

func (r *Response) WithCode(argCode code.Code) interface{} {
	r.SetCode(argCode)
	r.SetData(r.GetData())
	return r
}

func (r *Response) WithCodeAndMsg(argCode code.Code, msg string) interface{} {
	r.SetCode(argCode)
	r.SetMsg(msg)
	r.SetData(r.GetData())
	return r
}

func (r *Response) WithError(argCode code.Code, err error) interface{} {
	if argCode == 0 {
		argCode = code.StdErr
	}
	r.SetCode(argCode)
	r.SetMsg(err.Error())
	r.SetData(r.GetData())
	return r
}

func (r *Response) Json(ctx *gin.Context) {
	if r.GetMsg() == "" {
		r.SetMsg(r.GetCode().GetMsg())
	}
	ctx.JSON(http.StatusOK, r)
	return
}
