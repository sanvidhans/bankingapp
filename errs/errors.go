package errs

import "net/http"

type AppError struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e AppError ) AsMessage() *AppError  {
	return &AppError{Message: e.Message}
}


func NewNotFoundError(message string) *AppError {
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnExpectedError(message string) *AppError {
	return &AppError{
		Code: http.StatusInternalServerError,
		Message: message,
	}
}

func NewUnProcessableInputs(message string) *AppError {
	return &AppError{
		Code: http.StatusUnprocessableEntity,
		Message: message,
	}
}