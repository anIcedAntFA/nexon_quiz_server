package usertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	hasher "nexon_quiz/components/hasher"
	userbusiness "nexon_quiz/modules/user/business"
	userentity "nexon_quiz/modules/user/entity"
	userstorage "nexon_quiz/modules/user/storage"

	"github.com/gin-gonic/gin"
)

func HandleRegisterUser(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var newUser userentity.UserCreate

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := userstorage.NewUserMySQLStorage(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterUserBusiness(storage, md5)

		if err := business.Register(ctx.Request.Context(), &newUser); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(newUser.Id))
	}
}
