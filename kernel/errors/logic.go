package errors

import "gin.go.tpl/kernel/code"

type LogicError struct {
	BaseError
}

func (le LogicError) GenFromStdError(err error) Error {
	return LogicError{BaseError{code: code.StdErr, msg: err.Error(), cause: err}}
}

func (le LogicError) GenFromMsg(msg string, cause error) Error {
	return LogicError{BaseError{code: code.StdErr, msg: msg, cause: cause}}
}

func (le LogicError) GenFromCode(tmpCode code.Code, cause error) Error {
	return LogicError{BaseError{code: tmpCode, msg: tmpCode.GetMsg(), cause: cause}}
}

func (le LogicError) GenFromComplete(tmpCode code.Code, errMsg string, cause error) Error {
	return LogicError{BaseError{code: tmpCode, msg: errMsg, cause: cause}}
}

func (le LogicError) Unique(cause error) Error {
	return le.GenFromCode(code.StdDbUnique, cause)
}

func (le LogicError) NotFound(cause error) Error {
	return le.GenFromCode(code.StdDbNotFound, cause)
}
