package questiontransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	questionbusiness "nexon_quiz/modules/question/business"
	questionstorage "nexon_quiz/modules/question/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleGetQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		storage := questionstorage.NewQuestionMySQLStorage(db)
		business := questionbusiness.NewFindQuestionBusiness(storage)

		question, err := business.FindQuestion(
			ctx.Request.Context(),
			map[string]interface{}{"id": id},
		)

		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"Get question successfully",
			question,
		))
	}
}
