package categorybusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
)

type CreateCategoryStorage interface {
	InsertNewCategory(
		ctx context.Context,
		category *categoryentity.CategoryCreate,
	) error

	FindCategory(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*categoryentity.Category, error)
}

type createCategoryBusiness struct {
	storage CreateCategoryStorage
}

func NewCreateCategoryBusiness(storage CreateCategoryStorage) *createCategoryBusiness {
	return &createCategoryBusiness{
		storage: storage,
	}
}

func (biz *createCategoryBusiness) CreateNewCategory(
	ctx context.Context,
	newCategory *categoryentity.CategoryCreate,
) error {
	// currentCategory, err := biz.storage.FindCategory(ctx, map[string]interface{}{"content": })

	if err := newCategory.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	newCategory.Prepare(newCategory.DeletedAt)

	if err := biz.storage.InsertNewCategory(ctx, newCategory); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			categoryentity.ErrorCannotCreateCategory.Error(),
			err.Error(),
			"CannotCreateCategory",
		)
	}

	return nil
}
