package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionStorage interface {
	CreateQuestion(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
	) error
}

type createQuestionBusiness struct {
	storage CreateQuestionStorage
}

func NewCreateQuestionBusiness(storage CreateQuestionStorage) *createQuestionBusiness {
	return &createQuestionBusiness{
		storage: storage,
	}
}

func (biz *createQuestionBusiness) CreateQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
) error {
	if err := newQuestion.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.storage.CreateQuestion(ctx, newQuestion); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	return nil
}
