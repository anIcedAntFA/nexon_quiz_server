package gamesettingbusiness

import (
	"context"
	"nexon_quiz/common"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"

	"gorm.io/gorm"
)

type FindGameSettingStorage interface {
	FindGameSettingByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*gamesettingentity.GameSetting, error)
}

type findGameSettingBusiness struct {
	storage FindGameSettingStorage
}

func NewFindGameSettingBusiness(storage FindGameSettingStorage) *findGameSettingBusiness {
	return &findGameSettingBusiness{storage: storage}
}

func (biz *findGameSettingBusiness) GetGameSettingByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*gamesettingentity.GameSetting, error) {
	data, err := biz.storage.FindGameSettingByCondition(
		ctx,
		condition,
		"TypeSettings",
		"DifficultySetting",
		"CategorySettings",
	)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gamesettingentity.ErrorGameSettingNotFound
		}

		return nil, common.NewCustomError(
			err,
			gamesettingentity.ErrorGameSettingNotFound.Error(),
			"ErrorCannotGetGameSetting",
		)
	}

	return data, err
}
