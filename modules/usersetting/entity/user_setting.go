package usersettingentity

import (
	"nexon_quiz/common"

	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
)

type UserSetting struct {
	common.SQLModel
	UserId            uuid.UUID `json:"user_id" gorm:"column:user_id;"`
	Quantity          int       `json:"quantity" gorm:"column:quantity;"`
	TypeSettingId     uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
	DifficultyId      uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingId uuid.UUID `json:"category_setting_id" gorm:"column:category_setting_id;"`
	// TypeSetting       typesettingentity.TypeSetting         `json:"type_settings" gorm:"preload:false;"`
	// CategorySetting   categorysettingentity.CategorySetting `json:"category_settings" gorm:"preload:false;"`
	TypeSettings []typeentity.Type `json:"type_settings" gorm:"many2many:type_settings;"`
}

func (UserSetting) TableName() string {
	return "user_settings"
}
