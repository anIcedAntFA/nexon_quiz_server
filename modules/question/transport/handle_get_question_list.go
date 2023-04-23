package questiontransport

import (
	"log"
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	questionbusiness "nexon_quiz/modules/question/business"
	questionentity "nexon_quiz/modules/question/entity"
	questionstorage "nexon_quiz/modules/question/storage"

	"github.com/gin-gonic/gin"
)

func HandleGetQuestionList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryParams common.QueryParams

		if err := ctx.ShouldBindQuery(&queryParams); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		queryParams.Fulfill()

		var filter questionentity.Filter

		if err := ctx.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := questionstorage.NewQuestionMySQLStorage(db)

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)
		log.Println("user role", requester.GetRole())

		business := questionbusiness.NewQuestionListBusiness(storage, requester)

		data, pagination, err := business.QuestionList(ctx.Request.Context(), &filter, &queryParams)

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
