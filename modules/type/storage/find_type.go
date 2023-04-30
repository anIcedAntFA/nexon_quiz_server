package typestorage

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"gorm.io/gorm"
)

func (ts *typeMySQLStorage) FindTypeByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*typeentity.Type, error) {
	var data typeentity.Type

	if err := ts.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, typeentity.ErrorTypeNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
