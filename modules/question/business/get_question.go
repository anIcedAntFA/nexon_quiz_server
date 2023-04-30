package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"gorm.io/gorm"
)

type FindQuestionStorage interface {
	FindQuestionByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*questionentity.Question, error)
}

type findQuestionBusiness struct {
	storage FindQuestionStorage
}

func NewFindQuestionBusiness(storage FindQuestionStorage) *findQuestionBusiness {
	return &findQuestionBusiness{storage: storage}
}

func (biz *findQuestionBusiness) GetQuestionByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*questionentity.Question, error) {
	result, err := biz.storage.FindQuestionByCondition(ctx, condition, "Answers")

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, questionentity.ErrorQuestionNotFound
		}

		return nil, common.NewCustomError(
			err,
			questionentity.ErrorCannotGetQuestion.Error(),
			"ErrorCannotGetQuestion",
		)
	}

	return result, err
}
