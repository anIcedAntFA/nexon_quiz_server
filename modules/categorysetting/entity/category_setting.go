package usersettingentity

import (
	"nexon_quiz/common"
	categoryentity "nexon_quiz/modules/category/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategorySetting struct {
	common.SQLModel
	CategoryId            uuid.UUID                 `json:"category_id" gorm:"column:category_id;"`
	UserCategorySettingId uuid.UUID                 `json:"user_category_setting_id" gorm:"column:user_category_setting_id;"`
	Categories            []categoryentity.Category `json:"categories" gorm:"preload:false;"`
}

func (CategorySetting) TableName() string {
	return "category_settings"
}

func (csc *CategorySettingCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	csc.Id = id

	return err
}

type CategorySettingCreate struct {
	common.SQLModel
	CategoryId            uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	UserCategorySettingId uuid.UUID `json:"user_category_setting_id" gorm:"column:user_category_setting_id;"`
}

func (CategorySettingCreate) TableName() string {
	return CategorySetting{}.TableName()
}

type CategorySettingUpdate struct {
	common.SQLModel
	CategoryId            *uuid.UUID `json:"category_id" gorm:"column:category_id;"`
	UserCategorySettingId *uuid.UUID `json:"user_category_setting_id" gorm:"column:user_category_setting_id;"`
}
