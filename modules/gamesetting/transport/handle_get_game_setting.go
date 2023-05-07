package gamesettingtransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	gamesettingbusiness "nexon_quiz/modules/gamesetting/business"
	gamesettingstorage "nexon_quiz/modules/gamesetting/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetGameSettingById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := gamesettingstorage.NewGameSettingMySQLStorage(db)

		business := gamesettingbusiness.NewFindGameSettingBusiness(storage)

		data, err := business.GetGameSettingByCondition(
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
