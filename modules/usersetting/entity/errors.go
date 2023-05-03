package usersettingentity

import (
	"errors"
)

var (
	ErrorUserSettingIsBlank          = errors.New("user setting cannot be blank")
	ErrorUserSettingInvalid          = errors.New("user setting is not valid")
	ErrorUserSettingAlreadyExisted   = errors.New("user setting is already existed")
	ErrorUserSettingDeleted          = errors.New("user setting has been deleted")
	ErrorUserSettingNotFound         = errors.New("user setting not found")
	ErrorCannotCreateUserSetting     = errors.New("cannot create user setting")
	ErrorCannotCreateUserSettingList = errors.New("cannot create user setting list")
	ErrorCannotGetUserSetting        = errors.New("cannot get user setting")
	ErrorCannotGetListUserSetting    = errors.New("cannot get user setting list")
	ErrorCannotUpdateUserSetting     = errors.New("cannot update user setting")
	ErrorCannotDeleteUserSetting     = errors.New("cannot delete user setting")
	ErrorRequesterIsNotRootAdmin     = errors.New("no permission, only root admin can do this")
)
