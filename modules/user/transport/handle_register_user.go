package usertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	hasher "nexon_quiz/components/hasher"
	userbusiness "nexon_quiz/modules/user/business"
	userentity "nexon_quiz/modules/user/entity"
	userrepository "nexon_quiz/modules/user/repository"
	userstorage "nexon_quiz/modules/user/storage"
	userrolestorage "nexon_quiz/modules/userrole/storage"

	"github.com/gin-gonic/gin"
)

func HandleRegisterUser(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var newUser userentity.UserCreate

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		userStorage := userstorage.NewUserMySQLStorage(db)

		userRoleStorage := userrolestorage.NewUserRoleMySQLStorage(db)

		md5 := hasher.NewMd5Hash()

		repository := userrepository.NewRegisterUserRepository(userStorage, userRoleStorage, md5)

		business := userbusiness.NewRegisterUserBusiness(repository)

		if err := business.RegisterUser(ctx.Request.Context(), &newUser); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Register Successfully",
			newUser.Id,
		))
	}
}
