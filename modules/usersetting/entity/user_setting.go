package usersettingentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type UserSetting struct {
	common.SQLModel
	UserId            uuid.UUID `json:"user_id" gorm:"column:user_id;"`
	Quantity          int       `json:"quantity" gorm:"column:quantity;"`
	TypeSettingId     uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
	DifficultyId      uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingId uuid.UUID `json:"category_setting_id" gorm:"column:category_setting_id;"`
}

func (UserSetting) TableName() string {
	return "user_settings"
}
