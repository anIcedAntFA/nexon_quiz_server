package userrolebusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	userroleentity "nexon_quiz/modules/userrole/entity"
)

type CreateUserRoleStorage interface {
	FindUserRole(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*userroleentity.UserRole, error)

	InsertNewUserRole(
		ctx context.Context,
		userRole *userroleentity.UserRoleCreate,
	) error
}

type createUserRoleBusiness struct {
	storage CreateUserRoleStorage
}

func NewCreateUserRoleBusiness(storage CreateUserRoleStorage) *createUserRoleBusiness {
	return &createUserRoleBusiness{
		storage: storage,
	}
}

func (biz *createUserRoleBusiness) CreateNewUserRole(
	ctx context.Context,
	userRole *userroleentity.UserRoleCreate,
) error {
	// oldUserRole, err := biz.storage.FindUserRole(ctx, map[string]interface{}{"content": userRole.Content})
	// log.Println("oldUserRole", oldUserRole)

	// if oldUserRole != nil {
	// 	log.Println("userRole.Content", userRole.Content)
	// 	log.Println("oldUserRole.Content", oldUserRole.Content)
	// 	log.Println("err", err)

	// 	if userRole.Content == oldUserRole.Content {
	// 		log.Println("hello")

	// 		return common.ErrorInvalidRequest(err)
	// 	}
	// }

	if err := userRole.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	userRole.Prepare(userRole.DeletedAt)

	if err := biz.storage.InsertNewUserRole(ctx, userRole); err != nil {
		return common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			userroleentity.ErrorCannotCreateUserRole.Error(),
			err.Error(),
			"CannotCreateUserRole",
		)
	}

	return nil
}
