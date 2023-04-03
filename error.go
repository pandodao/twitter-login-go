package twitter

import (
	"errors"
	"fmt"
)

var (
	ErrTwitterServerError = errors.New("twitter is unavailable")
	ErrInvalidRequest     = errors.New("twitter request is invalid")
)

type ErrorWrapper struct {
	error  `json:"-"`
	Errors []struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	} `json:"errors"`
}

func (e *ErrorWrapper) Error() string {
	return fmt.Sprintf("twitter errors: %+v", e.Errors)
}
