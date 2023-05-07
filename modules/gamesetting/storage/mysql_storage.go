package gamesettingstorage

import "gorm.io/gorm"

type gameSettingMySQLStorage struct {
	db *gorm.DB
}

func NewGameSettingMySQLStorage(db *gorm.DB) *gameSettingMySQLStorage {
	return &gameSettingMySQLStorage{db: db}
}
