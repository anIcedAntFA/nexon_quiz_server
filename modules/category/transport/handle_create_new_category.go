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

func HandleCreateNewCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newCategory categoryentity.CategoryCreate

		if err := ctx.ShouldBindJSON(&newCategory); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := categorystorage.NewCategoryMySQLStorage(db)

		business := categorybusiness.NewCreateCategoryBusiness(storage)

		if err := business.CreateNewCategory(ctx.Request.Context(), &newCategory); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new user role successfully",
			newCategory.Id,
		))
	}
}
