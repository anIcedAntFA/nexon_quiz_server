package middleware

import (
	"errors"
	"fmt"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	"nexon_quiz/components/tokenprovider/jwt"
	userstorage "nexon_quiz/modules/user/storage"
	"strings"

	"github.com/gin-gonic/gin"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrorWrongAuthorizationHeader(nil)
	}

	return parts[1], nil
}

// 1. get token from header
// 2. validate token and parse to payload
// 3. from the token payload, we use user_id to find from DB

func RequiredAuthorization(appCtx appctx.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		storage := userstorage.NewUserMySQLStorage(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := storage.FindUser(ctx.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.IsDeleted == 1 {
			panic(common.ErrorNoPermission(errors.New("user has been deleted or banned")))
		}

		ctx.Set(common.CurrentUser, user)

		ctx.Next()
	}
}

func ErrorWrongAuthorizationHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authorization header"),
		fmt.Sprintf("ErrorWrongAuthorizationHeader"),
	)
}
