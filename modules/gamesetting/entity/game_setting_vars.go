package gamesettingentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type GameSettingCreateRequest struct {
	common.SQLModel
	UserId             uuid.UUID   `json:"-" gorm:"column:user_id;"`
	Quantity           int         `json:"quantity" gorm:"column:quantity;"`
	TypeSettingIds     []uuid.UUID `json:"type_setting_ids" gorm:"-"`
	DifficultyId       uuid.UUID   `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingIds []uuid.UUID `json:"category_setting_ids" gorm:"-"`
}

type GameSettingCreate struct {
	common.SQLModel
	UserId       uuid.UUID `json:"-" gorm:"column:user_id;"`
	Quantity     int       `json:"quantity" gorm:"column:quantity;"`
	DifficultyId uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
}

func (GameSettingCreate) TableName() string {
	return GameSetting{}.TableName()
}

func (usc *GameSettingCreate) Prepare(userId uuid.UUID) {
	usc.UserId = userId
}

type GameSettingUpdate struct {
	common.SQLModel
	UserId       *uuid.UUID `json:"-" gorm:"column:user_id;"`
	Quantity     *int       `json:"quantity" gorm:"column:quantity;"`
	DifficultyId *uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
}

func (GameSettingUpdate) TableName() string {
	return GameSetting{}.TableName()
}
