package questionstorage

import "gorm.io/gorm"

type questionMySQLStorage struct {
	db *gorm.DB
}

func NewQuestionMySQLStorage(db *gorm.DB) *questionMySQLStorage {
	return &questionMySQLStorage{db: db}
}
