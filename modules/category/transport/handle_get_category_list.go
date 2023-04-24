package categorytransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	categorybusiness "nexon_quiz/modules/category/business"
	categoryentity "nexon_quiz/modules/category/entity"
	categorystorage "nexon_quiz/modules/category/storage"

	"github.com/gin-gonic/gin"
)

func HandleGetQuestionList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryParams common.QueryParams

		if err := ctx.ShouldBindQuery(&queryParams); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		queryParams.Fulfill()

		var filter categoryentity.CategoryFilter

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := categorystorage.NewCategoryMySQLStorage(db)

		business := categorybusiness.NewFindCategoryListBusiness(storage)

		data, pagination, err := business.GetQuestionList(
			ctx.Request.Context(),
			&filter,
			&queryParams,
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.NewSuccessResponse(
			http.StatusOK,
			"Get question list successfully",
			data,
			pagination,
		))
	}
}
