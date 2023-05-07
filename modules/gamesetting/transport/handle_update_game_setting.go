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
	"github.com/google/uuid"
)

func HandleUpdateGameSetting(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := uuid.Parse(ctx.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var gamesettingRequest gamesettingentity.GameSettingCreateRequest

		if err := ctx.ShouldBindJSON(&gamesettingRequest); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()

		typeStorage := typestorage.NewTypeMySQLStorage(db)
		typeSettingStorage := typesettingstorage.NewTypeSettingMySQLStorage(db)
		categoryStorage := categorystorage.NewCategoryMySQLStorage(db)
		categorySettingStorage := categorysettingstorage.NewCategorySettingMySQLStorage(db)
		gameSettingStorage := gamesettingstorage.NewGameSettingMySQLStorage(db)

		repository := gamesettingrepository.NewUpdateGameSettingRepository(
			typeStorage,
			typeSettingStorage,
			categoryStorage,
			categorySettingStorage,
			gameSettingStorage,
		)

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		business := gamesettingbusiness.NewUpdateGameSettingBusiness(requester, repository)

		if err := business.UpdateNewGameSetting(
			ctx.Request.Context(),
			id,
			&gamesettingRequest,
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"update user setting successfully",
			true,
		))
	}
}
