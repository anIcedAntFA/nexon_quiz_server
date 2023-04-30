package typetransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	typebusiness "nexon_quiz/modules/type/business"
	typestorage "nexon_quiz/modules/type/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetTypeById(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := typestorage.NewTypeMySQLStorage(db)

		business := typebusiness.NewFindTypeBusiness(storage)

		data, err := business.GetTypeByCondition(
			ctx,
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"get question type successfully",
			data,
		))
	}
}
