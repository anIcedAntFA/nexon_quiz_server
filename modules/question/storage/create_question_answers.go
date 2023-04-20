package questionstorage

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

func (s *questionMySQLStorage) CreateQuestionAnswers(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
	// newAnswers answerentity.AnswersCreate,
) error {
	if err := s.db.Create(&newQuestion).Error; err != nil {
		return common.ErrorDB(err)
	}

	// if err := s.db.Create(&newQuestion.Answers).Error; err != nil {
	// 	return common.ErrorDB(err)
	// }

	return nil
}
