package answerstorage

import "gorm.io/gorm"

type answerMySQLStorage struct {
	db *gorm.DB
}

func NewAnswerMySQLStorage(db *gorm.DB) *answerMySQLStorage {
	return &answerMySQLStorage{
		db: db,
	}
}
