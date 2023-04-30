package typebusiness

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteTypeStorage interface {
	FindTypeByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*typeentity.Type, error)

	DeleteTypeById(
		ctx context.Context,
		id uuid.UUID,
	) error
}

type deleteTypeBusiness struct {
	storage DeleteTypeStorage
}

func NewDeleteTypeBusiness(storage DeleteTypeStorage) *deleteTypeBusiness {
	return &deleteTypeBusiness{storage: storage}
}

func (biz *deleteTypeBusiness) DeleteTypeById(
	ctx context.Context,
	id uuid.UUID,
) error {
	oldType, err := biz.storage.FindTypeByCondition(
		ctx,
		map[string]interface{}{"id": id},
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return typeentity.ErrorTypeNotFound
		}

		return common.NewCustomError(
			err,
			typeentity.ErrorTypeNotFound.Error(),
			"ErrorCannotGetType",
		)
	}

	if oldType.DeletedAt != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorTypeDeleted.Error(),
			"ErrorTypeDeleted",
		)
	}

	if err := biz.storage.DeleteTypeById(ctx, id); err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotDeleteType.Error(),
			"ErrorCannotDeleteType",
		)
	}

	return nil
}
