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
	storage   CreateQuestionStorage
	requester common.Requester
}

func NewCreateQuestionBusiness(
	storage CreateQuestionStorage,
	requester common.Requester,
) *createQuestionBusiness {
	return &createQuestionBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *createQuestionBusiness) CreateQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
) error {
	if err := newQuestion.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	newQuestion.Prepare(biz.requester.GetRoleId(), 5, 5, 40)

	// newQuestion.TypeId

	if err := biz.storage.CreateQuestion(ctx, newQuestion); err != nil {
		return common.NewCustomError(
			err,
			questionentity.ErrorCannotCreateQuestion.Error(),
			"ErrorCannotCreateQuestion",
		)
	}

	return nil
}
