package answertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	answerbusiness "nexon_quiz/modules/answer/business"
	answerentity "nexon_quiz/modules/answer/entity"
	answerstorage "nexon_quiz/modules/answer/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateAnswerList(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// questionId, err := uuid.Parse(ctx.Param("question_id"))
		// log.Println("questionId", questionId)

		// if err != nil {
		// 	panic(common.ErrorInvalidRequest(err))
		// }

		var newAnswers answerentity.AnswersCreate

		if err := ctx.ShouldBindJSON(&newAnswers); err != nil {
			panic(err)
		}

		// newAnswer.QuestionId = questionId

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)

		business := answerbusiness.NewCreateAnswerListBusiness(storage)

		if err := business.CreateAnswerList(ctx, newAnswers); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(
			http.StatusOK,
			"Create answer list successfully",
			true,
		))
	}
}
