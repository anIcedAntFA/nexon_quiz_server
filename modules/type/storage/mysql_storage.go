package typestorage

import "gorm.io/gorm"

type typeMySQLStorage struct {
	db *gorm.DB
}

func NewTypeMySQLStorage(db *gorm.DB) *typeMySQLStorage {
	return &typeMySQLStorage{db: db}
}
