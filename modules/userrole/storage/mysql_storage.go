package userrolestorage

import "gorm.io/gorm"

type userRoleMySQLStorage struct {
	db *gorm.DB
}

func NewUserRoleMySQLStorage(db *gorm.DB) *userRoleMySQLStorage {
	return &userRoleMySQLStorage{db: db}
}
