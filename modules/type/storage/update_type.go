package typestorage

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
)

func (ts *typeMySQLStorage) UpdateTypeById(
	ctx context.Context,
	id uuid.UUID,
	newType *typeentity.TypeUpdate,
) error {
	if err := ts.db.Table(typeentity.Type{}.TableName()).
		Where("id = ?", id).
		Updates(newType).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
