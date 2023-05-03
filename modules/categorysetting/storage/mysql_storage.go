package categorysettingstorage

import "gorm.io/gorm"

type categorySettingMySQLStorage struct {
	db *gorm.DB
}

func NewCategorySettingMySQLStorage(db *gorm.DB) *categorySettingMySQLStorage {
	return &categorySettingMySQLStorage{db: db}
}
