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

func HandleGetTypeList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryParams common.QueryParams

		if err := ctx.ShouldBindQuery(&queryParams); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		queryParams.Fulfill()

		var filter typeentity.TypeFilter

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := typestorage.NewTypeMySQLStorage(db)

		business := typebusiness.NewFindTypeListBusiness(storage)

		data, pagination, err := business.GetTypeList(
			ctx.Request.Context(),
			&filter,
			&queryParams,
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(
			http.StatusOK,
			"Get question type list successfully",
			data,
			pagination,
		))
	}
}
