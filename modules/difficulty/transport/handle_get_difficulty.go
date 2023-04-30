package typetransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	difficultybusiness "nexon_quiz/modules/difficulty/business"
	difficultystorage "nexon_quiz/modules/difficulty/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetDifficultyById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := difficultystorage.NewDifficultyMySQLStorage(db)

		business := difficultybusiness.NewFindDifficultyBusiness(storage)

		data, err := business.GetDifficultyByCondition(
			ctx,
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"get question difficulty successfully",
			data,
		))
	}
}
