package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) CreateQuestionList(
	ctx context.Context,
	newQuestion []questionentity.QuestionCreate,
) error {
	db := s.db

	if err := db.Create(&newQuestion).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
