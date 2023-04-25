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

func HandleCreateNewDifficulty(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newDifficulty difficultyentity.DifficultyCreate

		if err := ctx.ShouldBindJSON(&newDifficulty); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := difficultystorage.NewDifficultyMySQLStorage(db)

		business := difficultybusiness.NewCreateDifficultyBusiness(storage)

		if err := business.CreateNewDifficulty(ctx.Request.Context(), &newDifficulty); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new question difficulty successfully",
			newDifficulty.Id,
		))
	}
}
