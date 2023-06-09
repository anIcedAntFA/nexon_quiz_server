package answerbusiness

import (
	"context"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
)

type CreateAnswerStorage interface {
	CreateAnswer(ctx context.Context, newAnswer *answerentity.AnswerCreate) error
}

type createAnswerBusiness struct {
	storage CreateAnswerStorage
}

func NewCreateAnswerBusiness(storage CreateAnswerStorage) *createAnswerBusiness {
	return &createAnswerBusiness{storage: storage}
}

func (biz *createAnswerBusiness) CreateAnswer(
	ctx context.Context,
	newAnswer *answerentity.AnswerCreate,
) error {
	if err := newAnswer.Validate(); err != nil {
		return common.ErrorInvalidRequest(err)
	}

	if err := biz.storage.CreateAnswer(ctx, newAnswer); err != nil {
		return common.ErrorCannotCreateEntity(answerentity.EntityName, err)
	}

	return nil
}
