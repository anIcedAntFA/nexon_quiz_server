package gamesettingtransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	categorystorage "nexon_quiz/modules/category/storage"
	categorysettingstorage "nexon_quiz/modules/categorysetting/storage"
	gamesettingbusiness "nexon_quiz/modules/gamesetting/business"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	gamesettingrepository "nexon_quiz/modules/gamesetting/repository"
	gamesettingstorage "nexon_quiz/modules/gamesetting/storage"
	typestorage "nexon_quiz/modules/type/storage"
	typesettingstorage "nexon_quiz/modules/typesetting/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewGameSetting(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var gamesettingRequest gamesettingentity.GameSettingCreateRequest

		if err := ctx.ShouldBindJSON(&gamesettingRequest); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		typeStorage := typestorage.NewTypeMySQLStorage(db)
		typeSettingStorage := typesettingstorage.NewTypeSettingMySQLStorage(db)
		categoryStorage := categorystorage.NewCategoryMySQLStorage(db)
		categorySettingStorage := categorysettingstorage.NewCategorySettingMySQLStorage(db)
		gameSettingStorage := gamesettingstorage.NewGameSettingMySQLStorage(db)

		repository := gamesettingrepository.NewCreateGameSettingRepository(
			typeStorage,
			typeSettingStorage,
			categoryStorage,
			categorySettingStorage,
			gameSettingStorage,
		)

		business := gamesettingbusiness.NewCreateGameSettingBusiness(requester, repository)

		if err := business.CreateNewGameSetting(
			ctx.Request.Context(),
			&gamesettingRequest,
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"create user setting successfully",
			gamesettingRequest.Id,
		))
	}
}
