package questiontransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	questionbusiness "nexon_quiz/modules/question/business"
	questionentity "nexon_quiz/modules/question/entity"
	questionstorage "nexon_quiz/modules/question/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateQuestionList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newQuestions []questionentity.QuestionCreate

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		if err := ctx.ShouldBindJSON(&newQuestions); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		for _, v := range newQuestions {
			v.OwnerId = requester.GetUserId()
		}

		db := appCtx.GetMainDBConnection()

		storage := questionstorage.NewQuestionMySQLStorage(db)

		business := questionbusiness.NewCreateQuestionListBusiness(storage)

		if err := business.CreateQuestionList(ctx.Request.Context(), newQuestions); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new question successfully",
			true,
		))
	}
}
