package userbusiness

import (
	"context"
	"nexon_quiz/common"
	"nexon_quiz/components/tokenprovider"
	userentity "nexon_quiz/modules/user/entity"
)

type LoginStorage interface {
	FindUser(
		ctx context.Context,
		condition map[string]interface{},
		moreInfo ...string,
	) (*userentity.User, error)
}

type loginBusiness struct {
	storage       LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(
	storage LoginStorage,
	tokenProvider tokenprovider.Provider,
	hasher Hasher,
	expiry int,
) *loginBusiness {
	return &loginBusiness{
		storage:       storage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1. find user, email
// 2. hash password from input and compare with password in db
// 3. provider: issue JWT token for client
// 		access token and refresh token
// 4. return token(s)

func (biz *loginBusiness) Login(
	ctx context.Context,
	data *userentity.UserLogin,
) (*common.SimpleUser, *tokenprovider.Token, error) {
	user, err := biz.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	simpleUser := common.SimpleUser{
		SQLModel: user.SQLModel,
		Username: user.Username,
		RoleId:   user.RoleId,
	}

	if err != nil {
		return nil, nil, userentity.ErrorEmailOrPasswordInvalid
	}

	hashedPassword := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != hashedPassword {
		return nil, nil, userentity.ErrorEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId:   user.Id,
		RoleId:   user.RoleId,
		Username: user.Username,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)

	if err != nil {
		return nil, nil, common.ErrorInternal(err)
	}

	return &simpleUser, accessToken, nil
}
