package exception

import (
	"errors"
	"fmt"
)

type Exception struct {
	code     Code
	message  string
	error    error
	errorMap map[string]string
}

func (e *Exception) GetError() string {
	if e.error != nil {
		return e.error.Error()
	}
	return ""
}

func (e *Exception) GetErrorMap() map[string]string {
	if e.errorMap != nil {
		return e.errorMap
	}
	return nil
}

func (e *Exception) GetCode() Code {
	return e.code
}

func (e *Exception) GetMessage() string {
	return e.message
}

func (e *Exception) GetDetailError() error {
	return errors.New(fmt.Sprintf("msg: %s, error: %v", e.message, e.error))
}

func _createException(code Code, message string, err error, errorMap map[string]string) *Exception {
	return &Exception{
		code:     code,
		message:  message,
		error:    err,
		errorMap: errorMap,
	}
}
