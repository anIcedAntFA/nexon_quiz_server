package usersettingbusiness

import (
	"context"
	"nexon_quiz/common"
	usersettingentity "nexon_quiz/modules/usersetting/entity"
)

type CreateUserSettingRepository interface {
	CreateNewUserSetting(
		ctx context.Context,
		userSettingRequest *usersettingentity.UserSettingCreateRequest,
	) error
}

type createUserSettingBusiness struct {
	requester  common.Requester
	repository CreateUserSettingRepository
}

func NewCreateUserSettingBusiness(
	requester common.Requester,
	repository CreateUserSettingRepository,
) *createUserSettingBusiness {
	return &createUserSettingBusiness{
		requester:  requester,
		repository: repository,
	}
}

func (biz *createUserSettingBusiness) CreateNewUserSetting(
	ctx context.Context,
	userSettingRequest *usersettingentity.UserSettingCreateRequest,
) error {
	userSettingRequest.UserId = biz.requester.GetUserId()

	if err := biz.repository.CreateNewUserSetting(ctx, userSettingRequest); err != nil {
		return err
	}

	return nil
}
