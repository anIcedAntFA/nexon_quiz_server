package userbusiness

import (
	"context"
	"nexon_quiz/common"
	userentity "nexon_quiz/modules/user/entity"
)

type RegisterUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*userentity.User, error)

	CreateUser(ctx context.Context, newUser *userentity.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerUserBusiness struct {
	storage RegisterUserStorage
	hasher  Hasher
}

func NewRegisterUserBusiness(storage RegisterUserStorage, hasher Hasher) *registerUserBusiness {
	return &registerUserBusiness{
		storage: storage,
		hasher:  hasher,
	}
}

func (biz *registerUserBusiness) Register(ctx context.Context, newUser *userentity.UserCreate) error {
	user, _ := biz.storage.FindUser(ctx, map[string]interface{}{"email": newUser.Email})

	if user != nil {
		if user.IsDeleted == 0 {
			return userentity.ErrorUserDisabledOrBanned
		}

		return userentity.ErrorEmailExisted
	}

	salt := common.GenerateSalt(50)

	newUser.Password = biz.hasher.Hash(newUser.Password + salt)
	newUser.Salt = salt
	newUser.Role = userentity.RoleUser

	if err := biz.storage.CreateUser(ctx, newUser); err != nil {
		return common.ErrorCannotCreateEntity(userentity.EntityName, err)
	}

	return nil
}
