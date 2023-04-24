package categorystorage

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

func (cs *categoryMySQLStorage) InsertNewCategory(
	ctx context.Context,
	category *categoryentity.CategoryCreate,
) error {
	if err := cs.db.Create(category).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
