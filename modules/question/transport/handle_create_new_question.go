package questiontransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	categorystorage "nexon_quiz/modules/category/storage"
	difficultystorage "nexon_quiz/modules/difficulty/storage"
	questionbusiness "nexon_quiz/modules/question/business"
	questionentity "nexon_quiz/modules/question/entity"
	questionrepository "nexon_quiz/modules/question/repository"
	questionstorage "nexon_quiz/modules/question/storage"
	typestorage "nexon_quiz/modules/type/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewQuestion(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newQuestion questionentity.QuestionCreate

		if err := ctx.ShouldBindJSON(&newQuestion); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		questionStorage := questionstorage.NewQuestionMySQLStorage(db)
		typeStorage := typestorage.NewTypeMySQLStorage(db)
		difficultyStorage := difficultystorage.NewDifficultyMySQLStorage(db)
		categoryStorage := categorystorage.NewCategoryMySQLStorage(db)

		repository := questionrepository.NewCreateQuestionRepository(
			questionStorage,
			typeStorage,
			difficultyStorage,
			categoryStorage,
		)

		business := questionbusiness.NewCreateQuestionBusiness(requester, repository)

		if err := business.CreateNewQuestion(ctx.Request.Context(), &newQuestion); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Create new question successfully",
			newQuestion.Id,
		))
	}
}
