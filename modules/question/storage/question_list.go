package questionstorage

import (
	"context"
	"fmt"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) QuestionList(
	ctx context.Context,
	// filter *todomodel.Filter,
	queryParams *common.QueryParams,
	moreKeys ...string,
) ([]questionentity.Question, error) {
	//requester := ctxs.MustGet(common.CurrentUser).(common.Requester)
	db := s.db.Table(questionentity.Question{}.TableName())

	// if f := filter; f != nil {
	// 	if len(f.Status) > 0 {
	// 		db = db.Where("status in (?)", f.Status)
	// 	}
	// 	if len(f.Type) > 0 {
	// 		db = db.Where("type in (?)", f.Type)
	// 	}
	// 	if len(f.Category) > 0 {
	// 		db = db.Where("category in (?)", f.Category)
	// 	}
	// 	if len(f.Level) > 0 {
	// 		db = db.Where("level in (?)", f.Level)
	// 	}
	// 	if f.Score > 0 {
	// 		db = db.Where("level in (?)", f.Score)
	// 	}
	// }

	if err := db.Count(&queryParams.TotalItems).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	db = db.
		Preload("Answers", "is_deleted = 0")
	//for _, v := range moreKeys {
	//	db = db.Preload(v)
	//}

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

	var questions []questionentity.Question

	if err := db.
		Offset(offset).
		Limit(queryParams.PageSize).
		Order(order).
		Find(&questions).
		Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return questions, nil
}
