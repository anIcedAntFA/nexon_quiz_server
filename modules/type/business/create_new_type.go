package typebusiness

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"
)

type CreateTypeStorage interface {
	FindTypeByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*typeentity.Type, error)

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

	oldType, err := biz.storage.FindTypeByCondition(
		ctx,
		map[string]interface{}{"content": newType.Content},
	)

	if newType.Content == oldType.Content {
		return common.NewCustomError(
			err,
			typeentity.ErrorTypeAlreadyExisted.Error(),
			"ErrorTypeAlreadyExisted",
		)
	}

	newType.Prepare()

	if err := biz.storage.InsertNewType(ctx, newType); err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotCreateType.Error(),
			"CannotCreateCategory",
		)
	}

	return nil
}
