package typetransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	typebusiness "nexon_quiz/modules/type/business"
	typeentity "nexon_quiz/modules/type/entity"
	typestorage "nexon_quiz/modules/type/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewType(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newType typeentity.TypeCreate

		if err := ctx.ShouldBindJSON(&newType); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := typestorage.NewTypeMySQLStorage(db)

		business := typebusiness.NewCreateTypeBusiness(storage)

		if err := business.CreateNewType(ctx.Request.Context(), &newType); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new question type successfully",
			newType.Id,
		))
	}
}
