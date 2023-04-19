package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionAnswersRepository interface {
	CreateQuestionAnswers(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
		newAnswers answerentity.AnswersCreate,
	) error
}

type createQuestionAnswersBusiness struct {
	repository CreateQuestionAnswersRepository
}

func NewCreateQuestionAnswersBusiness(
	repository CreateQuestionAnswersRepository,
) *createQuestionAnswersBusiness {
	return &createQuestionAnswersBusiness{repository: repository}
}

func (biz *createQuestionAnswersBusiness) CreateQuestionAnswers(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
	newAnswers answerentity.AnswersCreate,
) error {
	if err := newQuestion.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.repository.CreateQuestionAnswers(ctx, newQuestion, newAnswers); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	return nil
}
