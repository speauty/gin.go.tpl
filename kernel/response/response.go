package response

import (
	"gin.go.tpl/kernel/code"
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
	// WithError 根据错误生成响应
	WithError(argCode code.Code, err error) IResponse
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

func (r *Response) WithError(argCode code.Code, err error) IResponse {
	//@todo should hook error inside this func or wrapping this response handle, like in the base controller
	if argCode == 0 {
		argCode = code.StdErr
	}
	r.SetCode(argCode)
	r.SetMsg(err.Error())
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
		r.SetMsg(r.GetCode().GetMsg())
	}
}
