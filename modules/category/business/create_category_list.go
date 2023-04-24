package categorybusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

type CreateCategoryListStorage interface {
	InsertCategoryList(
		ctx context.Context,
		newCategories []categoryentity.CategoryCreate,
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
	newCategories []categoryentity.CategoryCreate,
) error {
	for _, category := range newCategories {
		if err := category.Validate(); err != nil {
			return common.NewCustomError(
				err,
				err.Error(),
				"ErrorInvalidRequest",
			)
		}
	}

	if err := biz.storage.InsertCategoryList(ctx, newCategories); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			categoryentity.ErrorCannotCreateCategoryList.Error(),
			err.Error(),
			"CannotCreateCategoryList",
		)
	}

	return nil
}
