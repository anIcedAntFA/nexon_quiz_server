package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionStorage interface {
	CreateQuestion(ctx context.Context, newQuestion *questionentity.QuestionCreate, moreKeys ...string) error
}

// type CreateAnswersStorage interface {
// 	CreateAnswerList(
// 		ctx context.Context,
// 		newAnswers answerentity.AnswersCreate,
// 	) error
// }

type createQuestionBusiness struct {
	questionStorage CreateQuestionStorage
	// answersStorage  CreateAnswersStorage
}

func NewCreateQuestionBusiness(questionStorage CreateQuestionStorage) *createQuestionBusiness {
	return &createQuestionBusiness{
		questionStorage: questionStorage,
		// answersStorage:  answersStorage,
	}
}

func (biz *createQuestionBusiness) CreateQuestion(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
) error {
	if err := newQuestion.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.questionStorage.CreateQuestion(ctx, newQuestion, "Answers"); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	// for _, v := range newAnswers {
	// 	v.QuestionId = newQuestion.Id
	// }

	// if err := biz.answersStorage.CreateAnswerList(ctx, newAnswers); err != nil {
	// 	return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	// }

	return nil
}
