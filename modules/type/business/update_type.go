package typebusiness

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdateTypeStorage interface {
	FindTypeByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*typeentity.Type, error)

	UpdateTypeById(
		ctx context.Context,
		id uuid.UUID,
		newType *typeentity.TypeUpdate,
	) error
}

type updateTypeBusiness struct {
	storage UpdateTypeStorage
}

func NewUpdateTypeBusiness(storage UpdateTypeStorage) *updateTypeBusiness {
	return &updateTypeBusiness{storage: storage}
}

func (biz *updateTypeBusiness) UpdateTypeById(
	ctx context.Context,
	id uuid.UUID,
	newType *typeentity.TypeUpdate,
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

	if err := biz.storage.UpdateTypeById(ctx, id, newType); err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotUpdateType.Error(),
			"ErrorCannotUpdateType",
		)
	}

	return nil
}
