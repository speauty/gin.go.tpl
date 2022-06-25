package errors

import "gin.go.tpl/kernel/code"

func Core() *CoreError {
	return &CoreError{}
}

type CoreError struct {
	BaseError
}

func (e *CoreError) RouteNotFound() IError {
	return e.NewFromCode(code.StdRouteNotFound, nil)
}
