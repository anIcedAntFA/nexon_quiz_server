package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionRepository interface {
	CreateNewQuestion(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
	) error
}

type createQuestionBusiness struct {
	requester  common.Requester
	repository CreateQuestionRepository
}

func NewCreateQuestionBusiness(
	requester common.Requester,
	repository CreateQuestionRepository,
) *createQuestionBusiness {
	return &createQuestionBusiness{
		requester:  requester,
		repository: repository,
	}
}

func (biz *createQuestionBusiness) CreateNewQuestion(
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

	if err := biz.repository.CreateNewQuestion(ctx, newQuestion); err != nil {
		return err
	}

	return nil
}
