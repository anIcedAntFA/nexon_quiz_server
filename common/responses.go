package common

import (
	"errors"
	"net/http"
)

type successResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type pagingSuccessResponse struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"metadata,omitempty"`
}

func NewSuccessResponse(statusCode int, message string, data interface{}) *successResponse {
	return &successResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}

func NewPagingSuccessResponse(statusCode int, message string, data interface{}, pagination interface{}) *pagingSuccessResponse {
	return &pagingSuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}
}

func SimpleSuccessResponse(statusCode int, message string, data interface{}) *successResponse {
	return NewSuccessResponse(statusCode, message, data)
}

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}
