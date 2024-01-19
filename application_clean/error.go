package application_clean

import "go_error/common"

type Error interface {
	Error() string
	UserMessage() string
}

type DefaultError struct {
	err         *common.Error
	userMessage string
}

func (e *DefaultError) Error() string {
	return e.err.Error()
}

func (e *DefaultError) UserMessage() string {
	return e.userMessage
}
