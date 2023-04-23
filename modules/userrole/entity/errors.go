package userroleentity

import (
	"errors"
)

var (
	ErrorContentIsBlank          = errors.New("content cannot be blank")
	ErrorUserRoleInvalid         = errors.New("user role is not valid")
	ErrorUserRoleAlreadyExisted  = errors.New("user role is already existed")
	ErrorUserRoleDeleted         = errors.New("user role has been deleted")
	ErrorUserRoleNotFound        = errors.New("user role not found")
	ErrorCannotCreateUserRole    = errors.New("cannot create user role")
	ErrorCannotGetUserRole       = errors.New("cannot get user role")
	ErrorCannotGetListUserRole   = errors.New("cannot get list user roles")
	ErrorCannotUpdateUserRole    = errors.New("cannot update user role")
	ErrorCannotDeleteUserRole    = errors.New("cannot delete user role")
	ErrorRequesterIsNotRootAdmin = errors.New("no permission, only root admin can do this")
)
