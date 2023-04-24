package userroletransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	userrolebusiness "nexon_quiz/modules/userrole/business"
	userrolestorage "nexon_quiz/modules/userrole/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetUserRole(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := userrolestorage.NewUserRoleMySQLStorage(db)

		business := userrolebusiness.NewFindUserRoleBusiness(storage)

		data, err := business.GetUserRoleByCondition(
			ctx.Request.Context(),
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"Get user role successfully",
			data,
		))
	}
}
