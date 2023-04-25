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

func HandleCreateCategoryList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newCategories []categoryentity.CategoryCreate

		if err := ctx.ShouldBindJSON(&newCategories); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		questionStorage := categorystorage.NewCategoryMySQLStorage(db)

		business := categorybusiness.NewCreateCategoryListBusiness(questionStorage)

		if err := business.CreateCategoryList(ctx.Request.Context(), newCategories); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create category list successfully",
			true,
		))
	}
}
