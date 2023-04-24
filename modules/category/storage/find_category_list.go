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
	db := cs.db.Table(categoryentity.Category{}.TableName())

	if err := db.Count(&queryParams.TotalItems).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	if content := filter.Content; content != nil {
		db = db.Where("content = ?", *content)
	}

	var offset int

	var order string

	if qp := queryParams; qp != nil {
		searchStr := fmt.Sprintf("%%%s%%", qp.Search)

		if len(qp.Search) > 0 {
			db = db.Where("content LIKE ?", searchStr)
		}

		offset = (queryParams.CurrentPage - 1) * queryParams.PageSize

		order = fmt.Sprintf("%s %s", queryParams.SortBy, queryParams.OrderBy)
	}

	var data []categoryentity.Category

	if err := db.
		Offset(offset).
		Limit(queryParams.PageSize).
		Order(order).
		Find(&data).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
