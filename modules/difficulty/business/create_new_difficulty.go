package difficultybusiness

import (
	"context"
	"nexon_quiz/common"

	difficultyentity "nexon_quiz/modules/difficulty/entity"
)

type CreateDifficultyStorage interface {
	FindDifficultyByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*difficultyentity.Difficulty, error)

	InsertNewDifficulty(
		ctx context.Context,
		newDifficulty *difficultyentity.DifficultyCreate,
	) error
}

type createDifficultyBusiness struct {
	storage CreateDifficultyStorage
}

func NewCreateDifficultyBusiness(storage CreateDifficultyStorage) *createDifficultyBusiness {
	return &createDifficultyBusiness{
		storage: storage,
	}
}

func (biz *createDifficultyBusiness) CreateNewDifficulty(
	ctx context.Context,
	newDifficulty *difficultyentity.DifficultyCreate,
) error {
	if err := newDifficulty.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	oldDifficulty, err := biz.storage.FindDifficultyByCondition(
		ctx,
		map[string]interface{}{"content": newDifficulty.Content},
	)

	if err == nil && newDifficulty.Content == oldDifficulty.Content {
		return common.NewCustomError(
			err,
			difficultyentity.ErrorDifficultyAlreadyExisted.Error(),
			"ErrorDifficultyAlreadyExisted",
		)
	}

	newDifficulty.Prepare()

	if err := biz.storage.InsertNewDifficulty(ctx, newDifficulty); err != nil {
		return common.NewCustomError(
			err,
			difficultyentity.ErrorCannotCreateDifficulty.Error(),
			"CannotCreateCategory",
		)
	}

	return nil
}
