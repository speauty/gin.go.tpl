package errors

import "gin.go.tpl/lib/code"

type SysError struct {
	BaseError
}

func (se SysError) RouteNotFound() Error {
	return SysError{BaseError{}.GenFromCode(code.StdRouteNotFound, nil).(BaseError)}
}
