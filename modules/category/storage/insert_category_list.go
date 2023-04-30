package categorystorage

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

func (s *categoryMySQLStorage) InsertCategoryList(
	ctx context.Context,
	newCategories categoryentity.CategoriesCreate,
) error {
	db := s.db

	if err := db.Create(&newCategories).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
