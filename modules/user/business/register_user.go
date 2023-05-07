package userbusiness

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	userentity "nexon_quiz/modules/user/entity"
)

type RegisterUserRepository interface {
	RegisterUser(
		ctx context.Context,
		newUser *userentity.UserCreate,
		gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
	) error
}

type registerUserBusiness struct {
	repository RegisterUserRepository
}

func NewRegisterUserBusiness(repository RegisterUserRepository) *registerUserBusiness {
	return &registerUserBusiness{
		repository: repository,
	}
}

func (biz *registerUserBusiness) RegisterUser(
	ctx context.Context,
	newUser *userentity.UserCreate,
	gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
) error {
	if err := biz.repository.RegisterUser(ctx, newUser, gameSettingRequest); err != nil {
		return common.NewCustomError(
			err,
			userentity.ErrorCannotCreateUser.Error(),
			"ErrorCannotCreateUserRole",
		)
	}

	return nil
}
