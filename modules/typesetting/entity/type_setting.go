package usersettingentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeSetting struct {
	common.SQLModel
	TypeId        uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	TypeSettingId uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
}

func (TypeSetting) TableName() string {
	return "type_settings"
}

func (tsc *TypeSettingCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	tsc.Id = uuid.UUID(id)

	return err
}

type TypeSettingCreate struct {
	common.SQLModel
	TypeId        uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	TypeSettingId uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
}

func (TypeSettingCreate) TableName() string {
	return TypeSetting{}.TableName()
}

type TypeSettingUpdate struct {
	common.SQLModel
	TypeId        *uuid.UUID `json:"type_id" gorm:"column:type_id;"`
	TypeSettingId *uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
}
