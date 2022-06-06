package http

import "gin.go.tpl/lib/code"

type Response struct {
	Code code.Code   `json:"c"`
	Msg  string      `json:"m"`
	Data interface{} `json:"d"`
}

func (r *Response) GetCode() code.Code {
	if r == nil {
		return 0
	}
	return r.Code
}

func (r *Response) GetMsg() string {
	if r == nil {
		return ""
	}
	return r.Msg
}

func (r *Response) GetData() interface{} {
	if r == nil {
		return nil
	}
	return r.Data
}

func (r *Response) Default() *Response {
	return &Response{Code: code.StdOk, Msg: code.StdOk.GetMsg(), Data: nil}
}

func (r *Response) RespByCode(code code.Code) *Response {
	return &Response{Code: code, Msg: code.GetMsg(), Data: r.GetData()}
}

func (r *Response) RespByCodeAndMsg(code code.Code, msg string) *Response {
	return &Response{Code: code, Msg: msg, Data: r.GetData()}
}
