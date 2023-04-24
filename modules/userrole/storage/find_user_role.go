package userrolestorage

import (
	"context"
	"nexon_quiz/common"
	userroleentity "nexon_quiz/modules/userrole/entity"

	"gorm.io/gorm"
)

func (urs *userRoleMySQLStorage) FindUserRole(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*userroleentity.UserRole, error) {
	db := urs.db

	for _, v := range moreKeys {
		db = db.Preload(v)
	}

	var data userroleentity.UserRole

	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, userroleentity.ErrorUserRoleNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &data, nil
}
