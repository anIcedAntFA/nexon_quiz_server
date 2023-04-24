package categorybusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"gorm.io/gorm"
)

type FindCategoryStorage interface {
	FindCategory(
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
	data, err := biz.storage.FindCategory(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			categoryentity.ErrorCannotGetListCategory.Error(),
			err.Error(),
			"CannotGetCategory",
		)
	}

	return data, nil
}
