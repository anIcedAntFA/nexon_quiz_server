package categorystorage

import (
	"context"
	"fmt"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

func (cs *categoryMySQLStorage) FindCategoryList(
	ctx context.Context,
	filter *categoryentity.CategoryFilter,
	queryParams *common.QueryParams,
	moreKeys ...string,
) ([]categoryentity.Category, error) {
	db := cs.db.
		Table(categoryentity.Category{}.TableName()).
		Where("deleted_at IS NULL")

	if err := db.Count(&queryParams.TotalItems).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if content := filter.Content; content != nil {
		db = db.Where("content = ?", *content)
	}

	var order string

	if qp := queryParams; qp != nil {
		searchStr := fmt.Sprintf("%%%s%%", qp.Search)

		if len(qp.Search) > 0 {
			db = db.Where("content LIKE ?", searchStr)
		}

		order = fmt.Sprintf("%s %s", queryParams.SortBy, queryParams.OrderBy)
	}

	var data []categoryentity.Category

	if err := db.
		Select("*").
		Order(order).
		Find(&data).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
