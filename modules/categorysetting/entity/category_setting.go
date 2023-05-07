package usersettingentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type CategorySetting struct {
	common.SimpleSQLModel
	CategoryId    uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	GameSettingId uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}

func (CategorySetting) TableName() string {
	return "category_settings"
}

type CategorySettingCreate struct {
	common.SimpleSQLModel
	CategoryId    uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	GameSettingId uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}

func (CategorySettingCreate) TableName() string {
	return CategorySetting{}.TableName()
}

type CategorySettingUpdate struct {
	common.SimpleSQLModel
	CategoryId    *uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	GameSettingId *uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}
