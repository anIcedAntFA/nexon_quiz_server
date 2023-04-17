package userstorage

import "gorm.io/gorm"

type userMySQLStorage struct {
	db *gorm.DB
}

func NewUserMySQLStorage(db *gorm.DB) *userMySQLStorage {
	return &userMySQLStorage{
		db: db,
	}
}
