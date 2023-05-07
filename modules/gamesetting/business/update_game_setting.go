package gamesettingbusiness

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"

	"github.com/google/uuid"
)

type UpdateGameSettingRepository interface {
	UpdateGameSetting(
		ctx context.Context,
		id uuid.UUID,
		gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
	) error
}

type updateGameSettingBusiness struct {
	requester  common.Requester
	repository UpdateGameSettingRepository
}

func NewUpdateGameSettingBusiness(
	requester common.Requester,
	repository UpdateGameSettingRepository,
) *updateGameSettingBusiness {
	return &updateGameSettingBusiness{
		requester:  requester,
		repository: repository,
	}
}

func (biz *updateGameSettingBusiness) UpdateNewGameSetting(
	ctx context.Context,
	id uuid.UUID,
	gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
) error {
	if err := biz.repository.UpdateGameSetting(
		ctx,
		id,
		gameSettingRequest,
	); err != nil {
		return err
	}

	return nil
}
