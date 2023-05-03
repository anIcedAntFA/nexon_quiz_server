package typesettingstorage

import "gorm.io/gorm"

type typeSettingMySQLStorage struct {
	db *gorm.DB
}

func NewTypeSettingMySQLStorage(db *gorm.DB) *typeSettingMySQLStorage {
	return &typeSettingMySQLStorage{db: db}
}
