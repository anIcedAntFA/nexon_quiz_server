package categorystorage

import "gorm.io/gorm"

type categoryMySQLStorage struct {
	db *gorm.DB
}

func NewCategoryMySQLStorage(db *gorm.DB) *categoryMySQLStorage {
	return &categoryMySQLStorage{db: db}
}
