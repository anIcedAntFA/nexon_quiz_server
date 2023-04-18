package answerstorage

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
)

func (s *answerMySQLStorage) CreateAnswer(ctx context.Context, data *answerentity.AnswerCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
