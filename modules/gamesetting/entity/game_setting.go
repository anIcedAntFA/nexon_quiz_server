package gamesettingentity

import (
	"nexon_quiz/common"

	difficultyentity "nexon_quiz/modules/difficulty/entity"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
)

type GameSetting struct {
	common.SQLModel
	UserId       uuid.UUID `json:"user_id" gorm:"column:user_id;"`
	Quantity     int       `json:"quantity" gorm:"column:quantity;"`
	DifficultyId uuid.UUID `json:"-" gorm:"column:difficulty_id;"`
	// TypeSetting       typesettingentity.TypeSetting         `json:"type_settings" gorm:"preload:false;"`
	// CategorySetting   categorysettingentity.CategorySetting `json:"category_settings" gorm:"preload:false;"`
	TypeSettings      []typeentity.Type            `json:"type_settings" gorm:"many2many:type_settings;"`
	DifficultySetting *difficultyentity.Difficulty `json:"difficulty_setting" gorm:"foreignKey:DifficultyId;"`
}

func (GameSetting) TableName() string {
	return "game_settings"
}
