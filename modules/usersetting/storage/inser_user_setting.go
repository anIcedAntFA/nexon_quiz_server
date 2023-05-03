package usersettingstorage

import (
	"context"
	"nexon_quiz/common"
	usersettingentity "nexon_quiz/modules/usersetting/entity"
)

func (uss *userSettingMySQLStorage) InsertNewUserSetting(
	ctx context.Context,
	newUserSetting *usersettingentity.UserSettingCreate,
) error {
	if err := uss.db.Create(newUserSetting).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
