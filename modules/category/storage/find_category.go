package categorystorage

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"gorm.io/gorm"
)

func (cs *categoryMySQLStorage) FindCategoryByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*categoryentity.Category, error) {
	var data categoryentity.Category

	if err := cs.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, categoryentity.ErrorCategoryNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
