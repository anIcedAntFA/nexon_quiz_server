package userstorage

import (
	"context"
	"nexon_quiz/common"
	userentity "nexon_quiz/modules/user/entity"

	"gorm.io/gorm"
)

func (s *userMySQLStorage) FindUser(
	ctx context.Context,
	condition map[string]interface{},
	moreInfo ...string,
) (*userentity.User, error) {
	db := s.db.Table(userentity.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user userentity.User

	if err := db.Where(condition).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}

		return nil, common.ErrorDB(err)
	}

	return &user, nil
}
