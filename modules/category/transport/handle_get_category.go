package categorytransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	categorybusiness "nexon_quiz/modules/category/business"
	categorystorage "nexon_quiz/modules/category/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := categorystorage.NewCategoryMySQLStorage(db)

		business := categorybusiness.NewFindCategoryBusiness(storage)

		data, err := business.GetCategoryByCondition(
			ctx.Request.Context(),
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"Get user role successfully",
			data,
		))
	}
}
