package userrolebusiness

import (
	"context"
	"nexon_quiz/common"
	userroleentity "nexon_quiz/modules/userrole/entity"
)

type CreateUserRoleStorage interface {
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
	userRole.Prepare(userRole.DeletedAt)

	if err := userRole.Validate(); err != nil {
		return common.NewCustomError(
			err,
			err.Error(),
			"ErrorInvalidRequest",
		)
	}

	// if userRole.Content

	if err := biz.storage.InsertNewUserRole(ctx, userRole); err != nil {
		return common.NewCustomError(
			err,
			userroleentity.ErrorCannotCreateUserRole.Error(),
			"CannotCreateUserRole",
		)
	}

	return nil
}
