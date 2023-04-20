package questionrepository

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
	questionentity "nexon_quiz/modules/question/entity"
)

type CreateQuestionStorage interface {
	CreateQuestion(
		ctx context.Context,
		newQuestion *questionentity.QuestionCreate,
	) error
}

type CreateAnswersStorage interface {
	CreateAnswerList(
		ctx context.Context,
		newAnswers answerentity.AnswersCreate,
	) error
}

type createQuestionAnswersRepository struct {
	questionStorage CreateQuestionStorage
	answersStorage  CreateAnswersStorage
}

func NewCreateQuestionAnswersRepository(
	questionStorage CreateQuestionStorage,
	answersStorage CreateAnswersStorage,
) *createQuestionAnswersRepository {
	return &createQuestionAnswersRepository{
		questionStorage: questionStorage,
		answersStorage:  answersStorage,
	}
}

func (biz *createQuestionAnswersRepository) CreateQuestionAnswers(
	ctx context.Context,
	newQuestion *questionentity.QuestionCreate,
	newAnswers answerentity.AnswersCreate,
) error {
	if err := newQuestion.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.questionStorage.CreateQuestion(ctx, newQuestion); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	for _, v := range newAnswers {
		v.QuestionId = newQuestion.Id
	}

	if err := biz.answersStorage.CreateAnswerList(ctx, newAnswers); err != nil {
		return common.ErrorCannotCreateEntity(questionentity.EntityName, err)
	}

	return nil
}
