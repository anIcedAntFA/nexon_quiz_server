package typesettingbusiness

import (
	"context"
	"errors"
	"nexon_quiz/common"
	typesettingentity "nexon_quiz/modules/typesetting/entity"
)

type CreateTypeSettingStorage interface {
	InsertNewTypeSettingList(
		ctx context.Context,
		newTypeSettings []typesettingentity.TypeSettingCreate,
	) error
}

type createTypeSettingBusiness struct {
	requester common.Requester
	storage   CreateTypeSettingStorage
}

func NewCreateTypeSettingBusiness(
	requester common.Requester,
	storage CreateTypeSettingStorage,
) *createTypeSettingBusiness {
	return &createTypeSettingBusiness{
		requester: requester,
		storage:   storage,
	}
}

func (biz *createTypeSettingBusiness) CreateNewTypeSetting(
	ctx context.Context,
	newTypeSettings []typesettingentity.TypeSettingCreate,
) error {

	if err := biz.storage.InsertNewTypeSettingList(ctx, newTypeSettings); err != nil {
		return common.NewCustomError(
			err,
			errors.New("cannot create type setting").Error(),
			"ErrorCannotCreateUserSetting",
		)
	}

	return nil
}
