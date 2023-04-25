package difficultystorage

import "gorm.io/gorm"

type difficultyMySQLStorage struct {
	db *gorm.DB
}

func NewDifficultyMySQLStorage(db *gorm.DB) *difficultyMySQLStorage {
	return &difficultyMySQLStorage{db: db}
}
