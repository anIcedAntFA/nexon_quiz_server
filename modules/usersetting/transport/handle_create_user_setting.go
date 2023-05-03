package usersettingtransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"
	categorystorage "nexon_quiz/modules/category/storage"
	categorysettingstorage "nexon_quiz/modules/categorysetting/storage"
	typestorage "nexon_quiz/modules/type/storage"
	typesettingstorage "nexon_quiz/modules/typesetting/storage"
	usersettingbusiness "nexon_quiz/modules/usersetting/business"
	usersettingentity "nexon_quiz/modules/usersetting/entity"
	usersettingrepository "nexon_quiz/modules/usersetting/repository"
	usersettingstorage "nexon_quiz/modules/usersetting/storage"

	"github.com/gin-gonic/gin"
)

func HandleCreateNewUserSetting(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userSettingRequest usersettingentity.UserSettingCreateRequest

		if err := ctx.ShouldBindJSON(&userSettingRequest); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		requester := ctx.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()

		typeStorage := typestorage.NewTypeMySQLStorage(db)
		typeSettingStorage := typesettingstorage.NewTypeSettingMySQLStorage(db)
		categoryStorage := categorystorage.NewCategoryMySQLStorage(db)
		categorySettingStorage := categorysettingstorage.NewCategorySettingMySQLStorage(db)
		userSettingStorage := usersettingstorage.NewUserSettingMySQLStorage(db)

		repository := usersettingrepository.NewCreateUserSettingRepository(
			typeStorage,
			typeSettingStorage,
			categoryStorage,
			categorySettingStorage,
			userSettingStorage,
		)

		business := usersettingbusiness.NewCreateUserSettingBusiness(requester, repository)

		if err := business.CreateNewUserSetting(
			ctx.Request.Context(),
			&userSettingRequest,
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Save user setting successfully",
			userSettingRequest.Id,
		))
	}
}
