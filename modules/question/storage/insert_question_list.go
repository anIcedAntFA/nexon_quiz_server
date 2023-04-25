package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) InsertQuestionList(
	ctx context.Context,
	newQuestions []questionentity.QuestionCreate,
) error {
	db := s.db

	if err := db.Create(&newQuestions).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
