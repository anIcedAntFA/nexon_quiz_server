package difficultystorage

import (
	"context"
	"fmt"
	"nexon_quiz/common"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
)

func (ds *difficultyMySQLStorage) FindDifficultyList(
	ctx context.Context,
	filter *difficultyentity.DifficultyFilter,
	queryParams *common.QueryParams,
	moreKeys ...string,
) ([]difficultyentity.Difficulty, error) {
	db := ds.db.
		Table(difficultyentity.Difficulty{}.TableName()).
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

	var data []difficultyentity.Difficulty

	if err := db.
		Select("*").
		Order(order).
		Find(&data).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return data, nil
}
