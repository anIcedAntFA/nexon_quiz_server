package gamesettingentity

import (
	"errors"
)

var (
	ErrorGameSettingIsBlank          = errors.New("game setting cannot be blank")
	ErrorGameSettingInvalid          = errors.New("game setting is not valid")
	ErrorGameSettingAlreadyExisted   = errors.New("game setting is already existed")
	ErrorGameSettingDeleted          = errors.New("game setting has been deleted")
	ErrorGameSettingNotFound         = errors.New("game setting not found")
	ErrorCannotCreateGameSetting     = errors.New("cannot create game setting")
	ErrorCannotCreateGameSettingList = errors.New("cannot create game setting list")
	ErrorCannotGetGameSetting        = errors.New("cannot get game setting")
	ErrorCannotGetListGameSetting    = errors.New("cannot get game setting list")
	ErrorCannotUpdateGameSetting     = errors.New("cannot update game setting")
	ErrorCannotDeleteGameSetting     = errors.New("cannot delete game setting")
	ErrorRequesterIsNotRootAdmin     = errors.New("no permission, only root admin can do this")
)
