package userroletransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	userrolebusiness "nexon_quiz/modules/userrole/business"
	userroleentity "nexon_quiz/modules/userrole/entity"
	userrolestorage "nexon_quiz/modules/userrole/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewUserRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newUserRole userroleentity.UserRoleCreate

		if err := ctx.ShouldBindJSON(&newUserRole); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := userrolestorage.NewUserRoleMySQLStorage(db)

		business := userrolebusiness.NewCreateUserRoleBusiness(storage)

		if err := business.CreateNewUserRole(ctx.Request.Context(), &newUserRole); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new user role successfully",
			newUserRole.Id,
		))
	}
}
