package domain

import "go_error/common"

type Error interface {
	Error() string
}

type DefaultError struct {
	err *common.Error
}

func (e *DefaultError) Error() string {
	return e.err.Error()
}

func NewDefaultError(err *common.Error) Error {
	return &DefaultError{err: err}
}
