package utils

import (
	"fmt"
	"strings"
)

// CustomError ...
type CustomError struct {
	Message string
	Code    string
}

// NewCustomError ...
func NewCustomError(err error) CustomError {
	errorStr := err.Error()
	fmt.Println(errorStr)
	code := strings.Index(errorStr, "code = ")
	desc := strings.Index(errorStr, "desc = ")
	return CustomError{
		Message: errorStr[desc+len("desc = "):],
		Code:    errorStr[code+len("code = ") : desc-1],
	}
}
