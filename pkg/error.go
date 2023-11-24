package asaas

import (
	"fmt"
	"runtime"
)

type ErrorResponse struct {
	Code        string `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type Error *errorAsaas

type errorAsaas struct {
	Type ErrorType `json:"type,omitempty"`
	File string    `json:"file,omitempty"`
	Line int       `json:"line,omitempty"`
	Msg  string    `json:"err,omitempty"`
}

func NewError(typeError ErrorType, v ...any) Error {
	_, file, line, _ := runtime.Caller(1)
	return &errorAsaas{
		Type: typeError,
		Msg:  fmt.Sprint(v...),
		Line: line,
		File: file,
	}
}

func NewByErrorType(typeError ErrorType, err error) Error {
	if err == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(1)
	return &errorAsaas{
		Type: typeError,
		Msg:  err.Error(),
		Line: line,
		File: file,
	}
}

func NewByError(err error) Error {
	if err == nil {
		return nil
	}
	_, file, line, _ := runtime.Caller(1)
	return &errorAsaas{
		Type: ERROR_UNEXPECTED,
		Msg:  err.Error(),
		Line: line,
		File: file,
	}
}
