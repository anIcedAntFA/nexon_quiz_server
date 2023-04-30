package typebusiness

import (
	"context"
	"nexon_quiz/common"
	typeentity "nexon_quiz/modules/type/entity"

	"gorm.io/gorm"
)

type FindTypeStorage interface {
	FindTypeByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*typeentity.Type, error)
}

type findTypeBusiness struct {
	storage FindTypeStorage
}

func NewFindTypeBusiness(storage FindTypeStorage) *findTypeBusiness {
	return &findTypeBusiness{storage: storage}
}

func (biz *findTypeBusiness) GetTypeByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*typeentity.Type, error) {
	data, err := biz.storage.FindTypeByCondition(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, typeentity.ErrorTypeNotFound
		}

		return nil, common.NewCustomError(
			err,
			typeentity.ErrorTypeNotFound.Error(),
			"ErrorCannotGetType",
		)
	}

	return data, err
}
