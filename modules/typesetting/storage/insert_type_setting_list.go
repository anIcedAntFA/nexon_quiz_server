package typesettingstorage

import (
	"context"
	"nexon_quiz/common"
	typesettingentity "nexon_quiz/modules/typesetting/entity"
)

func (uss *typeSettingMySQLStorage) InsertNewTypeSettingList(
	ctx context.Context,
	newTypeSettings []typesettingentity.TypeSettingCreate,
) error {
	if err := uss.db.Create(newTypeSettings).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
