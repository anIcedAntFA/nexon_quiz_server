package typetransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	difficultybusiness "nexon_quiz/modules/difficulty/business"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
	difficultystorage "nexon_quiz/modules/difficulty/storage"

	"github.com/gin-gonic/gin"
)

func HandleGetDifficultyList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryParams common.QueryParams

		if err := ctx.ShouldBindQuery(&queryParams); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		queryParams.Fulfill()

		var filter difficultyentity.DifficultyFilter

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := difficultystorage.NewDifficultyMySQLStorage(db)

		business := difficultybusiness.NewFindDifficultyListBusiness(storage)

		data, pagination, err := business.GetDifficultyList(
			ctx.Request.Context(),
			&filter,
			&queryParams,
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(
			http.StatusOK,
			"Get question difficulty list successfully",
			data,
			pagination,
		))
	}
}
