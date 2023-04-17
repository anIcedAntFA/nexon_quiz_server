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

	CreateUser(ctx context.Context, data *userentity.UserCreate) error
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

func (biz *registerUserBusiness) Register(ctx context.Context, data *userentity.UserCreate) error {
	user, _ := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		if user.Status == 0 {
			return userentity.ErrorUserDisabledOrBanned
		}

		return userentity.ErrorEmailExisted
	}

	salt := common.GenerateSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = userentity.RoleUser

	if err := biz.storage.CreateUser(ctx, data); err != nil {
		return common.ErrorCannotCreateEntity(userentity.EntityName, err)
	}

	return nil
}
