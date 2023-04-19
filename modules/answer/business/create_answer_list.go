package answerbusiness

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
)

type CreateAnswerListStorage interface {
	CreateAnswerList(
		ctx context.Context,
		newAnswers answerentity.AnswersCreate,
	) error
}

type createAnswerListBusiness struct {
	storage CreateAnswerListStorage
}

func NewCreateAnswerListBusiness(storage CreateAnswerListStorage) *createAnswerListBusiness {
	return &createAnswerListBusiness{storage: storage}
}

func (biz *createAnswerListBusiness) CreateAnswerList(
	ctx context.Context,
	newAnswers answerentity.AnswersCreate,
) error {
	// if err := newAnswer.Validate(); err != nil {
	// 	return common.ErrorInvalidRequest(err)
	// }

	if err := biz.storage.CreateAnswerList(ctx, newAnswers); err != nil {
		return common.ErrorCannotCreateEntity(answerentity.EntityName, err)
	}

	return nil
}
