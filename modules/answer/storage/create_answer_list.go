package answerstorage

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
)

func (s *answerMySQLStorage) CreateAnswerList(
	ctx context.Context,
	newAnswers answerentity.AnswersCreate,
) error {
	if err := s.db.Create(&newAnswers).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
