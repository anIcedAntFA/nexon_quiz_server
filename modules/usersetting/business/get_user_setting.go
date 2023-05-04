package usersettingbusiness

import (
	"context"
	"nexon_quiz/common"
	usersettingentity "nexon_quiz/modules/usersetting/entity"

	"gorm.io/gorm"
)

type FindUserSettingStorage interface {
	FindUserSettingByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usersettingentity.UserSetting, error)
}

type findUserSettingBusiness struct {
	storage FindUserSettingStorage
}

func NewFindUserSettingBusiness(storage FindUserSettingStorage) *findUserSettingBusiness {
	return &findUserSettingBusiness{storage: storage}
}

func (biz *findUserSettingBusiness) GetUserSettingByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*usersettingentity.UserSetting, error) {
	data, err := biz.storage.FindUserSettingByCondition(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, usersettingentity.ErrorUserSettingNotFound
		}

		return nil, common.NewCustomError(
			err,
			usersettingentity.ErrorUserSettingNotFound.Error(),
			"ErrorCannotGetUserSetting",
		)
	}

	return data, err
}
