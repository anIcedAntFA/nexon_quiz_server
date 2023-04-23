package userrolestorage

import (
	"context"
	"nexon_quiz/common"
	userroleentity "nexon_quiz/modules/userrole/entity"
)

func (urs *userRoleMySQLStorage) InsertNewUserRole(
	ctx context.Context,
	userRole *userroleentity.UserRoleCreate,
) error {
	if err := urs.db.Create(userRole).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
