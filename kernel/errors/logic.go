package errors

func Logic() *LogicError {
	return &LogicError{}
}

type LogicError struct {
	BaseError
}
