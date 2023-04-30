package repository

import (
	"context"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionStorage interface {
	CreateQuestion(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
	) error
}
