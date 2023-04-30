package categorybusiness

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"gorm.io/gorm"
)

type FindCategoryStorage interface {
	FindCategoryByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*categoryentity.Category, error)
}

type findCategoryBusiness struct {
	storage FindCategoryStorage
}

func NewFindCategoryBusiness(storage FindCategoryStorage) *findCategoryBusiness {
	return &findCategoryBusiness{storage: storage}
}

func (biz *findCategoryBusiness) GetCategoryByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*categoryentity.Category, error) {
	data, err := biz.storage.FindCategoryByCondition(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, categoryentity.ErrorCategoryNotFound
		}

		return nil, common.NewCustomError(
			err,
			categoryentity.ErrorCategoryNotFound.Error(),
			"ErrorCannotGetCategory",
		)
	}

	return data, nil
}
