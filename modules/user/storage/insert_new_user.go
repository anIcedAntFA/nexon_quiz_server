package userstorage

import (
	"context"
	"nexon_quiz/common"
	userentity "nexon_quiz/modules/user/entity"
)

func (s *userMySQLStorage) InsertNewUser(
	ctx context.Context,
	newUser *userentity.UserCreate,
	moreKeys ...string,
) error {
	//if u implement many action => open transaction & commit
	db := s.db.Begin()

	db = db.Preload("UserRole")

	if err := db.Table(newUser.TableName()).Create(newUser).Error; err != nil {
		//if err, call rollback before, if not, connection will be stuck here
		//=> too many connection to DB => crash DB

		db.Rollback()

		return common.ErrorDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()

		return common.ErrorDB(err)
	}

	return nil
}
