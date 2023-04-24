package userrolebusiness

import (
	"context"
	"net/http"
	"nexon_quiz/common"
	userroleentity "nexon_quiz/modules/userrole/entity"

	"gorm.io/gorm"
)

type FindUserRoleStorage interface {
	FindUserRole(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*userroleentity.UserRole, error)
}

type findUserRoleBusiness struct {
	storage FindUserRoleStorage
}

func NewFindUserRoleBusiness(storage FindUserRoleStorage) *findUserRoleBusiness {
	return &findUserRoleBusiness{storage: storage}
}

func (biz *findUserRoleBusiness) GetUserRoleByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*userroleentity.UserRole, error) {
	data, err := biz.storage.FindUserRole(ctx, condition)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrorRecordNotFound
		}

		return nil, common.NewFullErrorResponse(
			http.StatusInternalServerError,
			err,
			userroleentity.ErrorCannotGetUserRole.Error(),
			err.Error(),
			"CannotGetUserRole",
		)
	}

	return data, nil
}
