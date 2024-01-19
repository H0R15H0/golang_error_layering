package common

import (
	"bytes"
	"fmt"
)

type ErrorCode int

func (c ErrorCode) String() string {
	switch c {
	case ErrorCodeUnknown:
		return "Unknown"
	case ErrorCodeInvalidArgument:
		return "InvalidArgument"
	default:
		return "Unknown"
	}
}

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeInvalidArgument
)

type Error struct {
	Op      string
	Code    ErrorCode
	Err     error
	Message string
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	if e.Op != "" {
		fmt.Fprintf(&buf, "%s: ", e.Op)
	}

	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		fmt.Fprintf(&buf, "<%s> ", e.Code.String())
		buf.WriteString(e.Message)
	}
	return buf.String()
}
