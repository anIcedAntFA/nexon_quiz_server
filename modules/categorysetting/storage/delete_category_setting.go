package categorysettingstorage

import (
	"context"
	"nexon_quiz/common"
	categorysettingentity "nexon_quiz/modules/categorysetting/entity"

	"github.com/google/uuid"
)

func (css *categorySettingMySQLStorage) DeleteCategorySettingById(
	ctx context.Context,
	id uuid.UUID,
) error {
	if err := css.db.
		Table(categorysettingentity.CategorySetting{}.TableName()).
		Delete(
			categorysettingentity.CategorySetting{}.TableName(),
			"game_setting_id = ?", id,
		).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
