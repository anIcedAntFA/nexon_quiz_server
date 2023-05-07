package gamesettingrepository

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
	categorysettingentity "nexon_quiz/modules/categorysetting/entity"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	typeentity "nexon_quiz/modules/type/entity"
	typesettingentity "nexon_quiz/modules/typesetting/entity"

	"github.com/google/uuid"
)

type UpdateGameSettingStorage interface {
	UpdateGameSettingById(
		ctx context.Context,
		id uuid.UUID,
		newGameSetting *gamesettingentity.GameSettingUpdate,
	) error
}

type UpdateTypeSettingStorage interface {
	DeleteTypeById(
		ctx context.Context,
		id uuid.UUID,
	) error

	InsertNewTypeSettingList(
		ctx context.Context,
		newTypeSettings []typesettingentity.TypeSettingCreate,
	) error
}

type UpdateCategorySettingStorage interface {
	DeleteCategorySettingById(
		ctx context.Context,
		id uuid.UUID,
	) error

	InsertNewCategorySettingList(
		ctx context.Context,
		newCategorySettings []categorysettingentity.CategorySettingCreate,
	) error
}

type updateGameSettingRepository struct {
	typeStorage            FindTypeStorage
	typeSettingStorage     UpdateTypeSettingStorage
	categoryStorage        FindCategoryStorage
	categorySettingStorage UpdateCategorySettingStorage
	gameSettingStorage     UpdateGameSettingStorage
}

func NewUpdateGameSettingRepository(
	typeStorage FindTypeStorage,
	typeSettingStorage UpdateTypeSettingStorage,
	categoryStorage FindCategoryStorage,
	categorySettingStorage UpdateCategorySettingStorage,
	gameSettingStorage UpdateGameSettingStorage,
) *updateGameSettingRepository {
	return &updateGameSettingRepository{
		typeStorage:            typeStorage,
		typeSettingStorage:     typeSettingStorage,
		categoryStorage:        categoryStorage,
		categorySettingStorage: categorySettingStorage,
		gameSettingStorage:     gameSettingStorage,
	}
}

func (repo *updateGameSettingRepository) UpdateGameSetting(
	ctx context.Context,
	id uuid.UUID,
	gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
) error {
	// Validate type setting request is valid
	oldTypes, err := repo.typeStorage.FindTypesByIds(
		ctx,
		gameSettingRequest.TypeSettingIds,
	)

	if len(oldTypes) != len(gameSettingRequest.TypeSettingIds) {
		return common.NewCustomError(
			err,
			typeentity.ErrorTypeInvalid.Error(),
			"ErrorTypeSettingInvalid",
		)
	}

	// Validate category setting request is valid
	oldCategories, err := repo.categoryStorage.FindCategoriesByIds(
		ctx,
		gameSettingRequest.CategorySettingIds,
	)

	if len(oldCategories) != len(gameSettingRequest.CategorySettingIds) {
		return common.NewCustomError(
			err,
			categoryentity.ErrorCategoryInvalid.Error(),
			"ErrorCategorySettingInvalid",
		)
	}

	// update game setting
	newGameSetting := gamesettingentity.GameSettingUpdate{
		Quantity:     &gameSettingRequest.Quantity,
		DifficultyId: &gameSettingRequest.DifficultyId,
	}

	if err := repo.gameSettingStorage.UpdateGameSettingById(
		ctx,
		id,
		&newGameSetting,
	); err != nil {
		return common.NewCustomError(
			err,
			gamesettingentity.ErrorCannotUpdateGameSetting.Error(),
			"ErrorCannotUpdateGameSetting",
		)
	}

	// delete type setting
	if err := repo.typeSettingStorage.DeleteTypeById(ctx, id); err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotDeleteType.Error(),
			"ErrorCannotDeleteTypeSetting",
		)
	}

	// create type setting
	newTypeSettings := make(
		[]typesettingentity.TypeSettingCreate,
		len(gameSettingRequest.TypeSettingIds),
	)

	for i := range newTypeSettings {
		newTypeSettings[i] = typesettingentity.TypeSettingCreate{
			TypeId:        gameSettingRequest.TypeSettingIds[i],
			GameSettingId: id,
		}
	}

	if err := repo.typeSettingStorage.InsertNewTypeSettingList(
		ctx,
		newTypeSettings,
	); err != nil {
		return common.NewCustomError(
			err,
			typeentity.ErrorCannotCreateType.Error(),
			"ErrorCannotCreateTypeSetting",
		)
	}

	// delete category setting
	if err := repo.categorySettingStorage.DeleteCategorySettingById(ctx, id); err != nil {
		return common.NewCustomError(
			err,
			categoryentity.ErrorCannotDeleteCategory.Error(),
			"ErrorCannotDeleteCategorySetting",
		)
	}

	//create category setting
	newCategorySettings := make(
		[]categorysettingentity.CategorySettingCreate,
		len(gameSettingRequest.CategorySettingIds),
	)

	for i := range newCategorySettings {
		newCategorySettings[i] = categorysettingentity.CategorySettingCreate{
			CategoryId:    gameSettingRequest.CategorySettingIds[i],
			GameSettingId: id,
		}
	}

	if err := repo.categorySettingStorage.InsertNewCategorySettingList(
		ctx,
		newCategorySettings,
	); err != nil {
		return common.NewCustomError(
			err,
			categoryentity.ErrorCannotCreateCategory.Error(),
			"ErrorCannotCreateCategorySetting",
		)
	}

	return nil
}
