package answerstorage

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
)

func (s *answerMySQLStorage) CreateAnswer(
	ctx context.Context,
	newAnswer *answerentity.AnswerCreate,
) error {
	if err := s.db.Create(&newAnswer).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
