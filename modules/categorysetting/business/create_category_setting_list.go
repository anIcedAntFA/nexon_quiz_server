package typesettingbusiness

import (
	"context"
	"errors"
	"nexon_quiz/common"
	categorysettingentity "nexon_quiz/modules/categorysetting/entity"
)

type CreateCategorySettingStorage interface {
	InsertNewCategorySettingList(
		ctx context.Context,
		newCategorySettings []categorysettingentity.CategorySettingCreate,
	) error
}

type createCategorySettingBusiness struct {
	requester common.Requester
	storage   CreateCategorySettingStorage
}

func NewCreateCategorySettingBusiness(
	requester common.Requester,
	storage CreateCategorySettingStorage,
) *createCategorySettingBusiness {
	return &createCategorySettingBusiness{
		requester: requester,
		storage:   storage,
	}
}

func (biz *createCategorySettingBusiness) CreateNewCategorySetting(
	ctx context.Context,
	newCategorySettings []categorysettingentity.CategorySettingCreate,
) error {

	if err := biz.storage.InsertNewCategorySettingList(ctx, newCategorySettings); err != nil {
		return common.NewCustomError(
			err,
			errors.New("cannot create type setting").Error(),
			"ErrorCannotCreateUserSetting",
		)
	}

	return nil
}
