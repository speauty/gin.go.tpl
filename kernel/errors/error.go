package errors

import "gin.go.tpl/kernel/code"

type Error interface {
	Error() string
	GetRealError() error
	GetCode() code.Code
}

type BaseError struct {
	code  code.Code
	msg   string
	cause error
}

func (be BaseError) Error() string {
	return be.msg
}

func (be BaseError) GetCode() code.Code {
	return be.code
}

func (be BaseError) GetRealError() error {
	return be.cause
}

func (be BaseError) GenFromStdError(err error) Error {
	return BaseError{code: code.StdErr, msg: err.Error(), cause: err}
}

func (be BaseError) GenFromMsg(msg string, cause error) Error {
	return BaseError{code: code.StdErr, msg: msg, cause: cause}
}

func (be BaseError) GenFromCode(tmpCode code.Code, cause error) Error {
	return BaseError{code: tmpCode, msg: tmpCode.GetMsg(), cause: cause}
}

func (be BaseError) GenFromComplete(tmpCode code.Code, errMsg string, cause error) Error {
	return BaseError{code: tmpCode, msg: errMsg, cause: cause}
}
