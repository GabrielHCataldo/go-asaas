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
	Err  string    `json:"err,omitempty"`
}

func NewError(typeError ErrorType, v ...any) Error {
	_, file, line, _ := runtime.Caller(1)
	return &errorAsaas{
		Type: typeError,
		Err:  fmt.Sprint(v...),
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
		Err:  err.Error(),
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
		Err:  err.Error(),
		Line: line,
		File: file,
	}
}
