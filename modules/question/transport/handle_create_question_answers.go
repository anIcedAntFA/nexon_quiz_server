package questiontransport

import (
	"log"
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	answerentity "nexon_quiz/modules/answer/entity"
	questionentity "nexon_quiz/modules/question/entity"

	"github.com/gin-gonic/gin"
)

func HandleCreateQuestionAnswers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newQuestion questionentity.QuestionCreate

		if err := ctx.ShouldBindJSON(&newQuestion); err != nil {
			panic(err)
		}

		log.Println("newQuestion", newQuestion)

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		newQuestion.OwnerId = requester.GetUserId()

		var newAnswers answerentity.AnswersCreate

		if err := ctx.ShouldBindJSON(&newAnswers); err != nil {
			panic(err)
		}

		log.Println("newAnswers", newAnswers)

		// db := appCtx.GetMainDBConnection()

		// questionStorage := questionstorage.NewQuestionMySQLStorage(db)
		// answersStorage := answerstorage.NewAnswerMySQLStorage(db)
		// repository := questionrepository.NewCreateQuestionAnswersRepository(questionStorage, answersStorage)
		// business := questionbusiness.NewCreateQuestionAnswersBusiness(repository)

		// if err := business.CreateQuestionAnswers(ctx.Request.Context(), &newQuestion, newAnswers); err != nil {
		// 	panic(err)
		// }

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create question successfully",
			newQuestion.Id,
		))
	}
}
