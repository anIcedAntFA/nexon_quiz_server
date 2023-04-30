package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/google/uuid"
)

type CreateQuestionListStorage interface {
	InsertQuestionList(
		ctx context.Context,
		newQuestions questionentity.QuestionsCreate,
	) error
}

type createQuestionListBusiness struct {
	storage   CreateQuestionListStorage
	requester common.Requester
}

func NewCreateQuestionListBusiness(
	storage CreateQuestionListStorage,
	requester common.Requester,
) *createQuestionListBusiness {
	return &createQuestionListBusiness{
		storage:   storage,
		requester: requester,
	}
}

func (biz *createQuestionListBusiness) CreateQuestionList(
	ctx context.Context,
	newQuestions questionentity.QuestionsCreate,
) error {
	for _, question := range newQuestions {
		if err := question.Validate(); err != nil {
			return common.NewCustomError(
				err,
				err.Error(),
				"ErrorInvalidRequest",
			)
		}

		question.Prepare(biz.requester.GetUserId(), 5, 5, 40)
	}

	for _, question := range newQuestions {
		id, _ := uuid.NewUUID()
		question.Id = uuid.UUID(id)

		for _, answer := range *question.Answers {
			answer.QuestionId = question.Id
		}
	}

	if err := biz.storage.InsertQuestionList(ctx, newQuestions); err != nil {
		return common.NewCustomError(
			err,
			questionentity.ErrorCannotCreateQuestionList.Error(),
			"ErrorCannotCreateQuestionList",
		)
	}

	return nil
}
