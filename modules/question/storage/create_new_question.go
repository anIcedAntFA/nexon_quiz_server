package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) CreateQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
	moreKeys ...string,
) error {
	db := s.db

	// for _, v := range moreKeys {
	// 	db = db.Preload(v)
	// }

	if err := db.Create(&newQuestion).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
