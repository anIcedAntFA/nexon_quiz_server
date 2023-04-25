package typestorage

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

func (ts *typeMySQLStorage) InsertNewType(
	ctx context.Context,
	newType *typeentity.TypeCreate,
) error {
	if err := ts.db.Create(newType).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
