package gamesettingstorage

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"

	"github.com/google/uuid"
)

func (gss *gameSettingMySQLStorage) UpdateGameSettingById(
	ctx context.Context,
	id uuid.UUID,
	newGameSetting *gamesettingentity.GameSettingUpdate,
) error {
	if err := gss.db.Table(gamesettingentity.GameSettingUpdate{}.TableName()).
		Where("id = ?", id).
		Updates(newGameSetting).
		Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
