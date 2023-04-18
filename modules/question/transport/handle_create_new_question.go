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

func HandleCreateNewQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()

		// requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		var newData questionentity.QuestionCreate

		if err := ctx.ShouldBind(&newData); err != nil {
			panic(err)
		}

		// newData.OwnerId = requester.GetUserId()

		storage := questionstorage.NewQuestionMySQLStorage(db)
		business := questionbusiness.NewCreateQuestionBusiness(storage)

		if err := business.CreateQuestion(ctx.Request.Context(), &newData); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newData.Id))
	}
}
