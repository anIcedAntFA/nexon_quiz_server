package typebusiness

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

type CreateTypeStorage interface {
	InsertNewType(
		ctx context.Context,
		newType *typeentity.TypeCreate,
	) (int64, error)
}

type createTypeBusiness struct {
	storage CreateTypeStorage
}

func NewCreateTypeBusiness(storage CreateTypeStorage) *createTypeBusiness {
	return &createTypeBusiness{storage: storage}
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

	rowsAffected, err := biz.storage.InsertNewType(ctx, newType)

	if rowsAffected < 1 {
		return common.NewCustomError(
			err,
			typeentity.ErrorTypeAlreadyExisted.Error(),
			"ErrorTypeAlreadyExisted",
		)
	}

	if err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotCreateType.Error(),
			"ErrorCannotCreateType",
		)
	}

	return nil
}
