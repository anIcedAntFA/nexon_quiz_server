package answertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	answerbusiness "nexon_quiz/modules/answer/business"
	answerentity "nexon_quiz/modules/answer/entity"
	answerstorage "nexon_quiz/modules/answer/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleCreateNewAnswer(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		questionId, err := uuid.Parse(ctx.Param("question_id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var newAnswer answerentity.AnswerCreate

		if err := ctx.ShouldBind(&newAnswer); err != nil {
			panic(err)
		}

		newAnswer.QuestionId = questionId

		db := appCtx.GetMainDBConnection()

		storage := answerstorage.NewAnswerMySQLStorage(db)
		business := answerbusiness.NewCreateAnswerBusiness(storage)

		if err := business.CreateAnswer(ctx, &newAnswer); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(newAnswer.Id))
	}
}
