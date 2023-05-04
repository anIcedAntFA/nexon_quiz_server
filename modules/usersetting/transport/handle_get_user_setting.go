package usersettingtransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	usersettingbusiness "nexon_quiz/modules/usersetting/business"
	usersettingstorage "nexon_quiz/modules/usersetting/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetUserSettingById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := usersettingstorage.NewUserSettingMySQLStorage(db)

		business := usersettingbusiness.NewFindUserSettingBusiness(storage)

		data, err := business.GetUserSettingByCondition(
			ctx,
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"get user setting successfully",
			data,
		))
	}
}
