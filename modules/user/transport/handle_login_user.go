package usertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	hasher "nexon_quiz/components/hasher"
	jwt "nexon_quiz/components/tokenprovider/jwt"
	userbusiness "nexon_quiz/modules/user/business"
	userentity "nexon_quiz/modules/user/entity"
	userstorage "nexon_quiz/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func HandleLoginUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data userentity.UserLogin

		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		storage := userstorage.NewUserMySQLStorage(db)

		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginBusiness(storage, tokenProvider, md5, 60*60*24*7) //7 days

		simpleUser, account, err := business.Login(ctx.Request.Context(), &data)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
			"message":     "Login successfully",
			"result": gin.H{
				"user":        simpleUser,
				"accessToken": account.Token,
			},
		})
	}
}
