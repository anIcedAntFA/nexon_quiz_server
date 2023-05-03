package usersettingstorage

import "gorm.io/gorm"

type userSettingMySQLStorage struct {
	db *gorm.DB
}

func NewUserSettingMySQLStorage(db *gorm.DB) *userSettingMySQLStorage {
	return &userSettingMySQLStorage{db: db}
}
