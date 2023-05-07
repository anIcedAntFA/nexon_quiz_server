package gamesettingstorage

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
)

func (gss *gameSettingMySQLStorage) InsertNewGameSetting(
	ctx context.Context,
	newGameSetting *gamesettingentity.GameSettingCreate,
) error {
	if err := gss.db.Create(newGameSetting).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
