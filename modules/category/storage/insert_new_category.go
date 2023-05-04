package categorystorage

import (
	"context"
	categoryentity "nexon_quiz/modules/category/entity"
)

func (cs *categoryMySQLStorage) InsertNewCategory(
	ctx context.Context,
	category *categoryentity.CategoryCreate,
) (int64, error) {
	result := cs.db.Where(categoryentity.CategoryCreate{
		Content: category.Content,
	}).FirstOrCreate(category)

	// if err != nil {
	// 	return common.ErrorDB(err)
	// }

	return result.RowsAffected, nil
}
