package categorystorage

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"gorm.io/gorm"
)

func (cs *categoryMySQLStorage) FindCategory(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*categoryentity.Category, error) {
	db := cs.db

	for _, v := range moreKeys {
		db = db.Preload(v)
	}

	var data categoryentity.Category

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, categoryentity.ErrorCategoryNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
