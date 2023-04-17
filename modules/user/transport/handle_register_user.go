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
		db := appCtx.GetMainDBConnection()
		var data userentity.UserCreate

		if err := ctx.ShouldBindJSON(&data); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		storage := userstorage.NewUserMySQLStorage(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewRegisterUserBusiness(storage, md5)

		if err := business.Register(ctx.Request.Context(), &data); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.Id))
	}
}
