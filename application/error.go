package application

type Error interface {
	Error() string
	UserMessage() string
}

type FieldError interface {
	Error() string
	ErrorCode() FieldErrorCode
}

type FieldErrorCode int
