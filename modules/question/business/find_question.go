package questionbusiness

import (
	"context"
	"nexon_quiz/common"
	questionentity "nexon_quiz/modules/question/entity"

	"gorm.io/gorm"
)

type FindQuestionStorage interface {
	FindQuestion(
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

func (biz findQuestionBusiness) FindQuestion(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*questionentity.Question, error) {
	result, err := biz.storage.FindQuestion(ctx, condition, "Answers")

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.ErrorEntityNotFound(questionentity.EntityName, err)
	}

	// var questionResponse questionentity.Question

	// questionResponse = questionentity.Question{
	// 	 Id
	// }

	return result, err
}
