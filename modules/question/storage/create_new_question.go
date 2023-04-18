package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) CreateQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
) error {
	if err := s.db.Create(&newQuestion).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
