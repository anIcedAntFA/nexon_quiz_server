package usertransport

import (
	"net/http"
	"nexon_quiz/common"
	"nexon_quiz/components/appctx"

	hasher "nexon_quiz/components/hasher"
	categorysettingstorage "nexon_quiz/modules/categorysetting/storage"
	gamesettingentity "nexon_quiz/modules/gamesetting/entity"
	gamesettingstorage "nexon_quiz/modules/gamesetting/storage"
	typesettingstorage "nexon_quiz/modules/typesetting/storage"
	userbusiness "nexon_quiz/modules/user/business"
	userentity "nexon_quiz/modules/user/entity"
	userrepository "nexon_quiz/modules/user/repository"
	userstorage "nexon_quiz/modules/user/storage"
	userrolestorage "nexon_quiz/modules/userrole/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleRegisterUser(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var newUser userentity.UserCreate

		if err := ctx.ShouldBindJSON(&newUser); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		gamesettingRequest := gamesettingentity.GameSettingCreateRequest{
			Quantity:           16,
			TypeSettingIds:     []uuid.UUID{uuid.MustParse("52621674-47cb-498d-9df6-b1a69d3dcc4a")},
			DifficultyId:       uuid.MustParse("6ae8d984-96c0-4e25-a0a5-3ae2ae4c1aeb"),
			CategorySettingIds: []uuid.UUID{uuid.MustParse("a98d022e-073a-45c2-9a8e-f4d0a8ea52db")},
		}

		db := appCtx.GetMainDBConnection()

		userStorage := userstorage.NewUserMySQLStorage(db)
		userRoleStorage := userrolestorage.NewUserRoleMySQLStorage(db)
		typeSettingStorage := typesettingstorage.NewTypeSettingMySQLStorage(db)
		categorySettingStorage := categorysettingstorage.NewCategorySettingMySQLStorage(db)
		gameSettingStorage := gamesettingstorage.NewGameSettingMySQLStorage(db)

		md5 := hasher.NewMd5Hash()

		repository := userrepository.NewRegisterUserRepository(
			userStorage,
			userRoleStorage,
			md5,
			typeSettingStorage,
			categorySettingStorage,
			gameSettingStorage,
		)

		business := userbusiness.NewRegisterUserBusiness(repository)

		if err := business.RegisterUser(
			ctx.Request.Context(),
			&newUser,
			&gamesettingRequest,
		); err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusCreated, common.SimpleSuccessResponse(
			http.StatusCreated,
			"Register Successfully",
			newUser.Id,
		))
	}
}
