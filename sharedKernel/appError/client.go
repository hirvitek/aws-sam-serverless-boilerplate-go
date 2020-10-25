package appError

import (
	"fmt"
)

type AppError interface {
	Error() string
	GetStatusCode() int
}

type appError struct {
	Operation string
	ErrType   string
	Err       error
	StatusCode int
}

func NewInternal(err error, operation string) *appError {
	return &appError{
		Operation: operation,
		ErrType:   internal,
		Err:       err,
		StatusCode: 500,
	}
}

func NewApi(err error, statusCode int) *appError {
	return &appError{
		ErrType:   api,
		Err:       err,
		StatusCode: statusCode,
	}
}

func (a appError) Error() string {
	return fmt.Sprintf("%v\n %v", a.Operation, a.Err)
}

func (a appError) GetStatusCode() int {
	return a.StatusCode
}
