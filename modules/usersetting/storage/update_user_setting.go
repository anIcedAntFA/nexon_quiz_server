package usersettingstorage

import (
	"context"
	"nexon_quiz/common"
	usersettingentity "nexon_quiz/modules/usersetting/entity"
)

func (uss *userSettingMySQLStorage) UpdateUserSettingById(
	ctx context.Context,
	newUserSetting *usersettingentity.UserSettingUpdate,
) error {
	if err := uss.db.Create(newUserSetting).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
