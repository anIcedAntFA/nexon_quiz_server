package typestorage

import (
	"context"
	"fmt"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

func (ts *typeMySQLStorage) FindTypeList(
	ctx context.Context,
	filter *typeentity.TypeFilter,
	queryParams *common.QueryParams,
	moreKeys ...string,
) ([]typeentity.Type, error) {
	db := ts.db.Table(typeentity.Type{}.TableName())

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

	var data []typeentity.Type

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
