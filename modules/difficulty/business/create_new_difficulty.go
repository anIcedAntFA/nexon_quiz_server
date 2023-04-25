package difficultybusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"

	difficultyentity "nexon_quiz/modules/difficulty/entity"
)

type CreateDifficultyStorage interface {
	InsertNewDifficulty(
		ctx context.Context,
		category *difficultyentity.DifficultyCreate,
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
	newType *difficultyentity.DifficultyCreate,
) error {
	if err := newType.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	newType.Prepare(newType.DeletedAt)

	if err := biz.storage.InsertNewDifficulty(ctx, newType); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			difficultyentity.ErrorCannotCreateDifficulty.Error(),
			err.Error(),
			"CannotCreateCategory",
		)
	}

	return nil
}
