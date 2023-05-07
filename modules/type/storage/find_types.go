package typestorage

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ts *typeMySQLStorage) FindTypesByIds(
	ctx context.Context,
	condition []uuid.UUID,
	moreKeys ...string,
) ([]typeentity.Type, error) {
	var data []typeentity.Type

	if err := ts.db.Where(condition).Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, typeentity.ErrorTypeNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return data, nil
}
