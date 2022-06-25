package response

import (
	"gin.go.tpl/kernel/code"
	"gin.go.tpl/kernel/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// New 生成新响应结构体 尽量调用该函数获取响应结构体, 相当于快捷方式之类的?
func New() IResponse {
	return &Response{}
}

type IResponse interface {
	// GetCode 获取响应代码
	GetCode() code.Code
	// GetMsg 获取提示信息
	GetMsg() string
	// GetData 获取响应数据
	GetData() interface{}
	// SetCode 设置响应代码
	SetCode(argCode code.Code) IResponse
	// SetMsg 设置提示信息
	SetMsg(msg string) IResponse
	// SetData 设置响应数据
	SetData(data interface{}) IResponse
	// Default 生成默认响应
	Default() interface{}
	// WithCode 根据代码生成响应
	WithCode(argCode code.Code) IResponse
	// WithCodeAndMsg 根据代码和提示生成响应
	WithCodeAndMsg(argCode code.Code, msg string) IResponse
	// WithIError 根据IError生成响应
	WithIError(err errors.IError) IResponse
	// Json Json响应类型
	Json(ctx *gin.Context)
	// JsonP JsonP响应类型(感觉和Json没啥差别呢)
	JsonP(ctx *gin.Context)
	// Xml Xml响应类型
	Xml(ctx *gin.Context)
}

type Response struct {
	Code code.Code   `json:"c"`
	Msg  string      `json:"m"`
	Data interface{} `json:"d"`
}

func (r *Response) GetCode() code.Code {
	return r.Code
}

func (r *Response) GetMsg() string {
	return r.Msg
}

func (r *Response) GetData() interface{} {
	return r.Data
}

func (r *Response) SetCode(argCode code.Code) IResponse {
	r.Code = argCode
	return r
}

func (r *Response) SetMsg(msg string) IResponse {
	r.Msg = msg
	return r
}

func (r *Response) SetData(data interface{}) IResponse {
	r.Data = data
	return r
}

func (r *Response) Default() interface{} {
	r.SetCode(code.StdOk)
	return r
}

func (r *Response) WithCode(argCode code.Code) IResponse {
	r.SetCode(argCode)
	r.SetData(r.GetData())
	return r
}

func (r *Response) WithCodeAndMsg(argCode code.Code, msg string) IResponse {
	r.SetCode(argCode)
	r.SetMsg(msg)
	r.SetData(r.GetData())
	return r
}

func (r *Response) WithIError(err errors.IError) IResponse {
	r.SetCode(err.GetCode())
	r.SetMsg(err.GetMsg())
	r.SetData(r.GetData())
	return r
}

func (r *Response) Json(ctx *gin.Context) {
	r.fillMsgWithCode()
	ctx.JSON(http.StatusOK, r)
	return
}

func (r *Response) JsonP(ctx *gin.Context) {
	r.fillMsgWithCode()
	ctx.JSONP(http.StatusOK, r)
	return
}

func (r *Response) Xml(ctx *gin.Context) {
	r.fillMsgWithCode()
	ctx.XML(http.StatusOK, r)
	return
}

func (r *Response) fillMsgWithCode() {
	if r.GetMsg() == "" {
		r.SetMsg(r.GetCode().Trans())
	}
}

func (r *Response) flush() {
	r.SetCode(0)
	r.SetMsg("")
	r.SetData(nil)
	return
}
