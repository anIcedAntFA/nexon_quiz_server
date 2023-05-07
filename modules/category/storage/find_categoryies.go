package categorystorage

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (cs *categoryMySQLStorage) FindCategoriesByIds(
	ctx context.Context,
	condition []uuid.UUID,
	moreKeys ...string,
) ([]categoryentity.Category, error) {
	var data []categoryentity.Category

	if err := cs.db.Where(condition).Find(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, categoryentity.ErrorCategoryNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return data, nil
}
