package gamesettingbusiness

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
)

type CreateGameSettingRepository interface {
	CreateNewGameSetting(
		ctx context.Context,
		gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
	) error
}

type createGameSettingBusiness struct {
	requester  common.Requester
	repository CreateGameSettingRepository
}

func NewCreateGameSettingBusiness(
	requester common.Requester,
	repository CreateGameSettingRepository,
) *createGameSettingBusiness {
	return &createGameSettingBusiness{
		requester:  requester,
		repository: repository,
	}
}

func (biz *createGameSettingBusiness) CreateNewGameSetting(
	ctx context.Context,
	gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
) error {
	gameSettingRequest.UserId = biz.requester.GetUserId()

	if err := biz.repository.CreateNewGameSetting(ctx, gameSettingRequest); err != nil {
		return err
	}

	return nil
}
