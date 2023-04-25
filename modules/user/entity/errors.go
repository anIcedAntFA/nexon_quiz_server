package userentity

import (
	"errors"
)

var (
	ErrorEmailOrPasswordInvalid = errors.New("email or password invalid")
	ErrorEmailExisted           = errors.New("email has already existed")
	ErrorUserDisabledOrBanned   = errors.New("user has been disabled or banned")
	ErrorUserInvalid            = errors.New("user is not valid")
	ErrorUserAlreadyExisted     = errors.New("user is already existed")
	ErrorUserDeleted            = errors.New("user has been deleted")
	ErrorUserNotFound           = errors.New("user not found")
	ErrorCannotCreateUser       = errors.New("cannot create user")
	ErrorCannotCreateUserList   = errors.New("cannot create user list")
	ErrorCannotGetUser          = errors.New("cannot get user")
	ErrorCannotGetListUser      = errors.New("cannot get user list")
	ErrorCannotUpdateUser       = errors.New("cannot update user")
	ErrorCannotDeleteUser       = errors.New("cannot delete user")
	ErrorRequesterIsNotAdmin    = errors.New("no permission, only owner can do this")
)
