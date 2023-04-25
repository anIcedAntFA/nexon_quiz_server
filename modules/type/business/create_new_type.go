package typebusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

type CreateTypeStorage interface {
	InsertNewType(
		ctx context.Context,
		category *typeentity.TypeCreate,
	) error
}

type createTypeBusiness struct {
	storage CreateTypeStorage
}

func NewCreateTypeBusiness(storage CreateTypeStorage) *createTypeBusiness {
	return &createTypeBusiness{
		storage: storage,
	}
}

func (biz *createTypeBusiness) CreateNewType(
	ctx context.Context,
	newType *typeentity.TypeCreate,
) error {
	if err := newType.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	newType.Prepare(newType.DeletedAt)

	if err := biz.storage.InsertNewType(ctx, newType); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			typeentity.ErrorCannotCreateType.Error(),
			err.Error(),
			"CannotCreateCategory",
		)
	}

	return nil
}
