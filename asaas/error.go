package asaas

import (
	"fmt"
	"runtime"
)

type Error *errorAsaas

type errorAsaas struct {
	Type ErrorType `json:"type,omitempty"`
	File string    `json:"file,omitempty"`
	Line int       `json:"line,omitempty"`
	Msg  string    `json:"err,omitempty"`
}

type ErrorResponse struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewError(typeError ErrorType, v ...any) Error {
	if typeError == ErrorTypeUnexpected {
		logErrorSkipCaller(5, v...)
	}
	_, file, line, _ := runtime.Caller(1)
	return &errorAsaas{
		Type: typeError,
		Msg:  fmt.Sprint(v...),
		Line: line,
		File: file,
	}
}

func NewByError(err error) Error {
	if err == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(1)
	logErrorSkipCaller(5, "error unexpected:", err)
	return &errorAsaas{
		Type: ErrorTypeUnexpected,
		Msg:  err.Error(),
		Line: line,
		File: file,
	}
}
