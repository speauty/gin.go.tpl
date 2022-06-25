package errors

import "gin.go.tpl/kernel/code"

type IError interface {
	error
	// GetCode 获取错误代码
	GetCode() code.Code
	// GetMsg 获取错误信息
	GetMsg() string
	// GetCause 获取实际错误
	GetCause() error
	// SetCode 设置错误代码
	SetCode(argCode code.Code) IError
	// SetMsg 设置错误信息
	SetMsg(msg string) IError
	// SetCause 设置实际错误
	SetCause(causeError error) IError
	// New 生成错误完全体
	New(argCode code.Code, errMsg string, cause error) IError
	// NewFromCode 根据错误代码生成错误
	NewFromCode(argCode code.Code, cause error) IError
	// NewFromMsg 根据错误信息生成错误
	NewFromMsg(msg string, cause error) IError
	// NewFromError 根据实际错误生成错误
	NewFromError(cause error) IError
}

type BaseError struct {
	code  code.Code
	msg   string
	cause error
}

func (e *BaseError) Error() string {
	if e.cause == nil {
		return e.msg
	}
	return e.cause.Error()
}

func (e *BaseError) GetCode() code.Code {
	return e.code
}

func (e *BaseError) GetMsg() string {
	return e.msg
}

func (e *BaseError) GetCause() error {
	return e
}

func (e *BaseError) SetCode(argCode code.Code) IError {
	e.code = argCode
	return e
}

func (e *BaseError) SetMsg(msg string) IError {
	e.msg = msg
	return e
}

func (e *BaseError) SetCause(causeError error) IError {
	e.cause = causeError
	return e
}

func (e *BaseError) New(argCode code.Code, errMsg string, cause error) IError {
	_ = e.SetCode(argCode).SetMsg(errMsg).SetCause(cause)
	e.fix()
	return e
}

func (e *BaseError) NewFromCode(argCode code.Code, cause error) IError {
	return e.New(argCode, "", cause)
}

func (e *BaseError) NewFromMsg(msg string, cause error) IError {
	return e.New(0, msg, cause)
}

func (e *BaseError) NewFromError(cause error) IError {
	return e.New(0, "", cause)
}

func (e *BaseError) fix() {
	if e.GetCode() == 0 { // 会重置ErrOK吗?
		_ = e.SetCode(code.StdErr)
	}

	if e.GetMsg() == "" {
		_ = e.SetMsg(e.GetCode().Trans())
	}
}
