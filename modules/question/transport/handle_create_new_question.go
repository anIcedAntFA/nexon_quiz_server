package questiontransport

import (
	"log"
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	answerbusiness "nexon_quiz/modules/answer/business"
	answerentity "nexon_quiz/modules/answer/entity"
	answerstorage "nexon_quiz/modules/answer/storage"
	questionbusiness "nexon_quiz/modules/question/business"
	questionentity "nexon_quiz/modules/question/entity"
	questionstorage "nexon_quiz/modules/question/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newQuestion questionentity.QuestionCreate

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		if err := ctx.ShouldBindJSON(&newQuestion); err != nil {
			panic(err)
		}

		newQuestion.OwnerId = requester.GetUserId()

		log.Println("newQuestion", newQuestion)

		var newAnswers answerentity.AnswersCreate

		log.Println("newAnswers", newAnswers)

		if err := ctx.ShouldBind(&newAnswers); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		questionStorage := questionstorage.NewQuestionMySQLStorage(db)
		business := questionbusiness.NewCreateQuestionBusiness(questionStorage)

		answersStorage := answerstorage.NewAnswerMySQLStorage(db)
		answerBusiness := answerbusiness.NewCreateAnswerListBusiness(answersStorage)

		if err := business.CreateQuestion(ctx.Request.Context(), &newQuestion); err != nil {
			panic(err)
		}

		if err := answerBusiness.CreateAnswerList(ctx, newAnswers); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new question successfully",
			newQuestion.Id,
		))
	}
}
