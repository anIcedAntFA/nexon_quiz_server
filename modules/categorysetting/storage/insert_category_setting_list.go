package categorysettingstorage

import (
	"context"
	"nexon_quiz/common"
	categorysettingentity "nexon_quiz/modules/categorysetting/entity"
)

func (uss *categorySettingMySQLStorage) InsertNewCategorySettingList(
	ctx context.Context,
	newCategorySettings []categorysettingentity.CategorySettingCreate,
) error {
	if err := uss.db.Create(newCategorySettings).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
