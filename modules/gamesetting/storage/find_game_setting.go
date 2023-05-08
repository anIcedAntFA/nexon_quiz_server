package gamesettingstorage

import (
	"context"
	"nexon_quiz/common"

	gamesettingentity "nexon_quiz/modules/gamesetting/entity"

	"gorm.io/gorm"
)

func (gss *gameSettingMySQLStorage) FindGameSettingByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*gamesettingentity.GameSetting, error) {
	db := gss.db

	for _, v := range moreKeys {
		db = db.Preload(v, "deleted_at IS NULL")
	}

	var data gamesettingentity.GameSetting

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gamesettingentity.ErrorGameSettingNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
