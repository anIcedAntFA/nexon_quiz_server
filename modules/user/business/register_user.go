package userbusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	userentity "nexon_quiz/modules/user/entity"
)

type RegisterUserStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*userentity.User, error)

	InsertNewUser(ctx context.Context, newUser *userentity.UserCreate) error
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

	newUser.Prepare(newUser.DeletedAt)

	// newUser.RoleId = "c39c2f6a-3ac9-4d6b-a21f-fd1ba94eec38"

	salt := common.GenerateSalt(50)

	newUser.Password = biz.hasher.Hash(newUser.Password + salt)

	newUser.Salt = salt

	if err := biz.storage.InsertNewUser(ctx, newUser); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			userentity.ErrorCannotCreateUser.Error(),
			err.Error(),
			"ErrorCannotCreateUserRole",
		)
	}

	return nil
}
