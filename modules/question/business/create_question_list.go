package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/google/uuid"
)

type CreateQuestionListStorage interface {
	CreateQuestionList(
		ctx context.Context,
		newQuestion []questionentity.QuestionCreate,
	) error
}

type createQuestionListBusiness struct {
	questionStorage CreateQuestionListStorage
}

func NewCreateQuestionListBusiness(
	questionStorage CreateQuestionListStorage,
) *createQuestionListBusiness {
	return &createQuestionListBusiness{
		questionStorage: questionStorage,
	}
}

func (biz *createQuestionListBusiness) CreateQuestionList(
	ctx context.Context,
	newQuestion []questionentity.QuestionCreate,
) error {
	// if err := newQuestion.Validate(); err != nil {
	// 	return common.ErrorInvalidRequest(err)
	// }

	for _, questionValue := range newQuestion {
		id, _ := uuid.NewUUID()
		questionValue.Id = uuid.UUID(id)

		for _, answerValue := range *questionValue.Answers {
			answerValue.QuestionId = questionValue.Id
		}
	}

	if err := biz.questionStorage.CreateQuestionList(ctx, newQuestion); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	return nil
}
