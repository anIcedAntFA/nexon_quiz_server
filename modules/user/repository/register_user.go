package userrepository

import (
	"context"
	"nexon_quiz/common"
	userentity "nexon_quiz/modules/user/entity"
	userroleentity "nexon_quiz/modules/userrole/entity"
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

type registerUserRepository struct {
	userStorage     RegisterUserStorage
	userRoleStorage UserRoleStorage
	hasher          Hasher
}

type Hasher interface {
	Hash(data string) string
}

func NewRegisterUserRepository(
	userStorage RegisterUserStorage,
	userRoleStorage UserRoleStorage,
	hasher Hasher,
) *registerUserRepository {
	return &registerUserRepository{
		userStorage:     userStorage,
		userRoleStorage: userRoleStorage,
		hasher:          hasher,
	}
}

func (repo *registerUserRepository) RegisterUser(
	ctx context.Context,
	newUser *userentity.UserCreate,
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

	return nil
}
