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

type FindTypeStorage interface {
	FindTypesByIds(
		ctx context.Context,
		condition []uuid.UUID,
		moreKeys ...string,
	) ([]typeentity.Type, error)
}

type CreateTypeSettingStorage interface {
	InsertNewTypeSettingList(
		ctx context.Context,
		newTypeSettings []typesettingentity.TypeSettingCreate,
	) error
}

type FindCategoryStorage interface {
	FindCategoriesByIds(
		ctx context.Context,
		condition []uuid.UUID,
		moreKeys ...string,
	) ([]categoryentity.Category, error)
}

type CreateCategorySettingStorage interface {
	InsertNewCategorySettingList(
		ctx context.Context,
		newCategorySettings []categorysettingentity.CategorySettingCreate,
	) error
}

type CreateGameSettingStorage interface {
	InsertNewGameSetting(
		ctx context.Context,
		newGameSetting *gamesettingentity.GameSettingCreate,
	) error
}

type createGameSettingRepository struct {
	typeStorage            FindTypeStorage
	typeSettingStorage     CreateTypeSettingStorage
	categoryStorage        FindCategoryStorage
	categorySettingStorage CreateCategorySettingStorage
	gameSettingStorage     CreateGameSettingStorage
}

func NewCreateGameSettingRepository(
	typeStorage FindTypeStorage,
	typeSettingStorage CreateTypeSettingStorage,
	categoryStorage FindCategoryStorage,
	categorySettingStorage CreateCategorySettingStorage,
	gameSettingStorage CreateGameSettingStorage,
) *createGameSettingRepository {
	return &createGameSettingRepository{
		typeStorage:            typeStorage,
		typeSettingStorage:     typeSettingStorage,
		categoryStorage:        categoryStorage,
		categorySettingStorage: categorySettingStorage,
		gameSettingStorage:     gameSettingStorage,
	}
}

func (repo *createGameSettingRepository) CreateNewGameSetting(
	ctx context.Context,
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

	// create game setting
	gameSettingId, _ := uuid.NewUUID()

	gameSettingRequest.Id = gameSettingId

	newGameSetting := gamesettingentity.GameSettingCreate{
		SQLModel:     common.NewSQLModel(gameSettingId),
		UserId:       gameSettingRequest.UserId,
		Quantity:     gameSettingRequest.Quantity,
		DifficultyId: gameSettingRequest.DifficultyId,
	}

	if err := repo.gameSettingStorage.InsertNewGameSetting(ctx, &newGameSetting); err != nil {
		return common.NewCustomError(
			err,
			gamesettingentity.ErrorCannotCreateGameSetting.Error(),
			"ErrorCannotCreateGameSetting",
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
			GameSettingId: newGameSetting.Id,
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

	// create category setting
	newCategorySettings := make(
		[]categorysettingentity.CategorySettingCreate,
		len(gameSettingRequest.TypeSettingIds),
	)

	for i := range newCategorySettings {
		newCategorySettings[i] = categorysettingentity.CategorySettingCreate{
			CategoryId:    gameSettingRequest.CategorySettingIds[i],
			GameSettingId: newGameSetting.Id,
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
