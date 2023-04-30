package typetransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	typebusiness "nexon_quiz/modules/type/business"
	typeentity "nexon_quiz/modules/type/entity"
	typestorage "nexon_quiz/modules/type/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleUpdateTypeById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var newType typeentity.TypeUpdate

		if err := ctx.ShouldBindJSON(&newType); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := typestorage.NewTypeMySQLStorage(db)

		business := typebusiness.NewUpdateTypeBusiness(storage)

		if err := business.UpdateTypeById(ctx, id, &newType); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"update question type successfully",
			true,
		))
	}
}
