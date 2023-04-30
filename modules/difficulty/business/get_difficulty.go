package difficultybusiness

import (
	"context"
	"nexon_quiz/common"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
	typeentity "nexon_quiz/modules/type/entity"

	"gorm.io/gorm"
)

type FindDifficultyStorage interface {
	FindDifficultyByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*difficultyentity.Difficulty, error)
}

type findDifficultyBusiness struct {
	storage FindDifficultyStorage
}

func NewFindDifficultyBusiness(
	storage FindDifficultyStorage,
) *findDifficultyBusiness {
	return &findDifficultyBusiness{storage: storage}
}

func (biz *findDifficultyBusiness) GetDifficultyByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*difficultyentity.Difficulty, error) {
	data, err := biz.storage.FindDifficultyByCondition(ctx, condition)

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
