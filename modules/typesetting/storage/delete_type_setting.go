package typesettingstorage

import (
	"context"
	"nexon_quiz/common"
	typesettingentity "nexon_quiz/modules/typesetting/entity"

	"github.com/google/uuid"
)

func (tss *typeSettingMySQLStorage) DeleteTypeById(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := tss.db.
		Table(typesettingentity.TypeSetting{}.TableName()).
		Delete(
			typesettingentity.TypeSetting{}.TableName(),
			"game_setting_id = ?", id,
		).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
