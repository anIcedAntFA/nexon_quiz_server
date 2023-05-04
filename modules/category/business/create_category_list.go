package categorybusiness

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

type CreateCategoryListStorage interface {
	FindCategoryByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*categoryentity.Category, error)

	InsertCategoryList(
		ctx context.Context,
		newCategories categoryentity.CategoriesCreate,
	) error
}

type createCategoryListBusiness struct {
	storage CreateCategoryListStorage
}

func NewCreateCategoryListBusiness(
	storage CreateCategoryListStorage,
) *createCategoryListBusiness {
	return &createCategoryListBusiness{
		storage: storage,
	}
}

func (biz *createCategoryListBusiness) CreateCategoryList(
	ctx context.Context,
	newCategories categoryentity.CategoriesCreate,
) error {
	for _, category := range newCategories {
		if err := category.Validate(); err != nil {
			return common.NewCustomError(
				err,
				err.Error(),
				"ErrorInvalidRequest",
			)
		}

		oldType, err := biz.storage.FindCategoryByCondition(
			ctx,
			map[string]interface{}{"content": category.Content},
		)

		if err == nil && category.Content == oldType.Content {
			return common.NewCustomError(
				err,
				categoryentity.ErrorCategoryAlreadyExisted.Error(),
				"ErrorCategoryAlreadyExisted",
			)
		}
	}

	if err := biz.storage.InsertCategoryList(ctx, newCategories); err != nil {
		return common.NewCustomError(
			err,
			categoryentity.ErrorCannotCreateCategoryList.Error(),
			"CannotCreateCategoryList",
		)
	}

	return nil
}
