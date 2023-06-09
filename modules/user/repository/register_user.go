package userrepository

import (
	"context"
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"
	categorysettingentity "nexon_quiz/modules/categorysetting/entity"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	typeentity "nexon_quiz/modules/type/entity"
	typesettingentity "nexon_quiz/modules/typesetting/entity"
	userentity "nexon_quiz/modules/user/entity"
	userroleentity "nexon_quiz/modules/userrole/entity"

	"github.com/google/uuid"
)

type RegisterUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*userentity.User, error)

	InsertNewUser(ctx context.Context, newUser *userentity.UserCreate) error
}

type UserRoleStorage interface {
	FindUserRole(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*userroleentity.UserRole, error)
}

type CreateTypeSettingStorage interface {
	InsertNewTypeSettingList(
		ctx context.Context,
		newTypeSettings []typesettingentity.TypeSettingCreate,
	) error
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

type registerUserRepository struct {
	userStorage            RegisterUserStorage
	userRoleStorage        UserRoleStorage
	hasher                 Hasher
	typeSettingStorage     CreateTypeSettingStorage
	categorySettingStorage CreateCategorySettingStorage
	gameSettingStorage     CreateGameSettingStorage
}

type Hasher interface {
	Hash(data string) string
}

func NewRegisterUserRepository(
	userStorage RegisterUserStorage,
	userRoleStorage UserRoleStorage,
	hasher Hasher,
	typeSettingStorage CreateTypeSettingStorage,
	categorySettingStorage CreateCategorySettingStorage,
	gameSettingStorage CreateGameSettingStorage,
) *registerUserRepository {
	return &registerUserRepository{
		userStorage:            userStorage,
		userRoleStorage:        userRoleStorage,
		hasher:                 hasher,
		typeSettingStorage:     typeSettingStorage,
		categorySettingStorage: categorySettingStorage,
		gameSettingStorage:     gameSettingStorage,
	}
}

func (repo *registerUserRepository) RegisterUser(
	ctx context.Context,
	newUser *userentity.UserCreate,
	gameSettingRequest *gamesettingentity.GameSettingCreateRequest,
) error {
	user, _ := repo.userStorage.FindUser(
		ctx,
		map[string]interface{}{"email": newUser.Email},
	)

	if user != nil {
		if user.DeletedAt != nil {
			return userentity.ErrorUserDisabledOrBanned
		}

		return userentity.ErrorEmailExisted
	}

	if err := newUser.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	userRole, _ := repo.userRoleStorage.FindUserRole(
		ctx,
		map[string]interface{}{"content": 2},
	)

	newUser.RoleId = userRole.Id

	salt := common.GenerateSalt(50)

	newUser.Password = repo.hasher.Hash(newUser.Password + salt)

	newUser.Salt = salt

	if err := repo.userStorage.InsertNewUser(ctx, newUser); err != nil {
		return common.NewCustomError(
			err,
			userentity.ErrorCannotCreateUser.Error(),
			"ErrorCannotCreateUserRole",
		)
	}

	// create game setting
	gameSettingId, _ := uuid.NewUUID()

	gameSettingRequest.Id = gameSettingId

	newGameSetting := gamesettingentity.GameSettingCreate{
		SQLModel:     common.NewSQLModel(gameSettingId),
		UserId:       newUser.Id,
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
