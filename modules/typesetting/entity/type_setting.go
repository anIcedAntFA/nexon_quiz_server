package usersettingentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type TypeSetting struct {
	common.SimpleSQLModel
	TypeId        uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	GameSettingId uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}

func (TypeSetting) TableName() string {
	return "type_settings"
}

type TypeSettingCreate struct {
	common.SimpleSQLModel
	TypeId        uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	GameSettingId uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}

func (TypeSettingCreate) TableName() string {
	return TypeSetting{}.TableName()
}

type TypeSettingUpdate struct {
	common.SimpleSQLModel
	TypeId        *uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	GameSettingId *uuid.UUID `json:"game_setting_id" gorm:"column:game_setting_id;"`
}
