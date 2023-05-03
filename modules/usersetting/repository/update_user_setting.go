package usersettingrepository

// import (
// 	"context"
// 	"nexon_quiz/common"
// 	categoryentity "nexon_quiz/modules/category/entity"
// 	categorysettingentity "nexon_quiz/modules/categorysetting/entity"
// 	typeentity "nexon_quiz/modules/type/entity"
// 	typesettingentity "nexon_quiz/modules/typesetting/entity"
// 	usersettingentity "nexon_quiz/modules/usersetting/entity"

// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// type FindTypeStorage interface {
// 	FindTypeByCondition(
// 		ctx context.Context,
// 		condition map[string]interface{},
// 		moreKeys ...string,
// 	) (*typeentity.Type, error)
// }

// type CreateTypeSettingStorage interface {
// 	InsertNewTypeSettingList(
// 		ctx context.Context,
// 		newTypeSettings []typesettingentity.TypeSettingCreate,
// 	) error
// }

// type FindCategoryStorage interface {
// 	FindCategoryByCondition(
// 		ctx context.Context,
// 		condition map[string]interface{},
// 		moreKeys ...string,
// 	) (*categoryentity.Category, error)
// }

// type CreateCategorySettingStorage interface {
// 	InsertNewCategorySettingList(
// 		ctx context.Context,
// 		newCategorySettings []categorysettingentity.CategorySettingCreate,
// 	) error
// }

// type UpdateUserSettingStorage interface {
// 	UpdateUserSettingById(
// 		ctx context.Context,
// 		newUserSetting *usersettingentity.UserSettingUpdate,
// 	) error
// }

// type createUserSettingRepository struct {
// 	typeStorage            FindTypeStorage
// 	typeSettingStorage     CreateTypeSettingStorage
// 	categoryStorage        FindCategoryStorage
// 	categorySettingStorage CreateCategorySettingStorage
// 	userSettingStorage     CreateUserSettingStorage
// }

// func NewCreateUserSettingRepository(
// 	typeStorage FindTypeStorage,
// 	typeSettingStorage CreateTypeSettingStorage,
// 	categoryStorage FindCategoryStorage,
// 	categorySettingStorage CreateCategorySettingStorage,
// 	userSettingStorage CreateUserSettingStorage,
// ) *createUserSettingRepository {
// 	return &createUserSettingRepository{
// 		typeStorage:            typeStorage,
// 		typeSettingStorage:     typeSettingStorage,
// 		categoryStorage:        categoryStorage,
// 		categorySettingStorage: categorySettingStorage,
// 		userSettingStorage:     userSettingStorage,
// 	}
// }

// func (repo *createUserSettingRepository) CreateNewUserSetting(
// 	ctx context.Context,
// 	userSettingRequest *usersettingentity.UserSettingCreateRequest,
// ) error {
// 	for _, typeId := range userSettingRequest.TypeSettingIds {
// 		_, err := repo.typeStorage.FindTypeByCondition(
// 			ctx,
// 			map[string]interface{}{"id": typeId},
// 		)

// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return typeentity.ErrorTypeNotFound
// 			}

// 			return common.NewCustomError(
// 				err,
// 				typeentity.ErrorTypeInvalid.Error(),
// 				"ErrorTypeSettingInvalid",
// 			)
// 		}
// 	}

// 	newTypeSettings := make(
// 		[]typesettingentity.TypeSettingCreate,
// 		len(userSettingRequest.TypeSettingIds),
// 	)

// 	typeSettingId, _ := uuid.NewUUID()

// 	for i := range newTypeSettings {
// 		newTypeSettings[i] = typesettingentity.TypeSettingCreate{
// 			TypeId:        userSettingRequest.TypeSettingIds[i],
// 			TypeSettingId: typeSettingId,
// 		}
// 	}

// 	for _, categoryId := range userSettingRequest.CategorySettingIds {
// 		_, err := repo.categoryStorage.FindCategoryByCondition(
// 			ctx,
// 			map[string]interface{}{"id": categoryId},
// 		)

// 		if err != nil {
// 			if err == gorm.ErrRecordNotFound {
// 				return categoryentity.ErrorCategoryNotFound
// 			}

// 			return common.NewCustomError(
// 				err,
// 				categoryentity.ErrorCategoryInvalid.Error(),
// 				"ErrorCategorySettingInvalid",
// 			)
// 		}
// 	}

// 	newCategoriesSettings := make(
// 		[]categorysettingentity.CategorySettingCreate,
// 		len(userSettingRequest.CategorySettingIds),
// 	)

// 	categorySettingId, _ := uuid.NewUUID()

// 	for i := range newCategoriesSettings {
// 		newCategoriesSettings[i] = categorysettingentity.CategorySettingCreate{
// 			CategoryId:        userSettingRequest.CategorySettingIds[i],
// 			CategorySettingId: categorySettingId,
// 		}
// 	}

// 	if err := repo.typeSettingStorage.InsertNewTypeSettingList(
// 		ctx,
// 		newTypeSettings,
// 	); err != nil {
// 		return common.NewCustomError(
// 			err,
// 			typeentity.ErrorCannotCreateType.Error(),
// 			"ErrorCannotCreateTypeSetting",
// 		)
// 	}

// 	if err := repo.categorySettingStorage.InsertNewCategorySettingList(
// 		ctx,
// 		newCategoriesSettings,
// 	); err != nil {
// 		return common.NewCustomError(
// 			err,
// 			categoryentity.ErrorCannotCreateCategory.Error(),
// 			"ErrorCannotCreateCategorySetting",
// 		)
// 	}

// 	newUserSetting := usersettingentity.UserSettingCreate{
// 		UserId:            userSettingRequest.UserId,
// 		Quantity:          userSettingRequest.Quantity,
// 		TypeSettingId:     typeSettingId,
// 		DifficultyId:      userSettingRequest.DifficultyId,
// 		CategorySettingId: categorySettingId,
// 	}

// 	if err := repo.userSettingStorage.InsertNewUserSetting(ctx, &newUserSetting); err != nil {
// 		return common.NewCustomError(
// 			err,
// 			usersettingentity.ErrorCannotCreateUserSetting.Error(),
// 			"ErrorCannotCreateUserSetting",
// 		)
// 	}

// 	return nil
// }
