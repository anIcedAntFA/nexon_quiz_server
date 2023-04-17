package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func ErrorDB(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong with DB",
		err.Error(),
		"DB_ERROR",
	)
}

func ErrorInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrorInvalidRequest")
}

func ErrorInternal(err error) *AppError {
	return NewFullErrorResponse(
		http.StatusInternalServerError,
		err,
		"something went wrong in the server",
		err.Error(),
		"ErrorInternal",
	)
}

func ErrorEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Error%sNotFound", entity),
	)
}

func ErrorCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotCreate%s", entity),
	)
}

func ErrorCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotGet%s", entity),
	)
}

func ErrorCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotList%s", entity),
	)
}

func ErrorCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotList%s", entity),
	)
}

func ErrorEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("User already exists %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorUserAlreadyExists%s", entity),
	)
}

func ErrorEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already deleted", strings.ToLower(entity)),
		fmt.Sprintf("Error%sDeleted", entity),
	)
}

func ErrorCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrorCannotDelete%s", entity),
	)
}

func ErrorNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("you have no permission"),
		fmt.Sprintf("ErrorNoPermission"),
	)
}

var ErrorRecordNotFound = errors.New("record not found")
