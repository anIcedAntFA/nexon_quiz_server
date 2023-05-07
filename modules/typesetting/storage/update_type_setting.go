package typesettingstorage

import (
	"context"
	"nexon_quiz/common"
	typesettingentity "nexon_quiz/modules/typesetting/entity"

	"github.com/google/uuid"
)

func (tss *typeSettingMySQLStorage) UpdateTypeSettingById(
	ctx context.Context,
	id uuid.UUID,
	newTypeSettings typesettingentity.TypeSettingUpdate,
) error {
	if err := tss.db.Table(typesettingentity.TypeSetting{}.TableName()).
		Where("game_setting_id = ?", id).
		Updates(newTypeSettings).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
