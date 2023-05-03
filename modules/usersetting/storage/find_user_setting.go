package usersettingstorage

import (
	"context"
	"nexon_quiz/common"

	usersettingentity "nexon_quiz/modules/usersetting/entity"

	"gorm.io/gorm"
)

func (uss *userSettingMySQLStorage) FindUserSettingByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usersettingentity.UserSetting, error) {
	db := uss.db

	for _, v := range moreKeys {
		db = db.Preload(v, "deleted_at IS NULL")
	}

	var data usersettingentity.UserSetting

	if err := uss.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, usersettingentity.ErrorUserSettingNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
