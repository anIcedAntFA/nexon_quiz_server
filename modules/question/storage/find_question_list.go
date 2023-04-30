package questionstorage

import (
	"context"
	"fmt"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) FindQuestionList(
	ctx context.Context,
	filter *questionentity.QuestionFilter,
	queryParams *common.QueryParams,
	moreKeys ...string,
) ([]questionentity.Question, error) {
	// requester := ctx.Value(common.CurrentUser).(common.Requester)

	db := s.db.
		Table(questionentity.Question{}.TableName()).
		Where("deleted_at IS NULL")

	if err := db.Count(&queryParams.TotalItems).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	db = db.
		Preload("Answers")

	if f := filter; f != nil {
		if f.Type != nil {
			db = db.Where("type in (?)", f.Type)
		}

		if f.Difficulty != "" {
			db = db.Where("difficulty = ?", f.Difficulty)
		}

		if f.Category != nil {
			db = db.Where("category in (?)", f.Category)
		}
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

	var data []questionentity.Question

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
