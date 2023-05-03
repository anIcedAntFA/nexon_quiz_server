package usersettingentity

import (
	"nexon_quiz/common"

	categoryentity "nexon_quiz/modules/category/entity"
	difficultyentity "nexon_quiz/modules/difficulty/entity"
	typeentity "nexon_quiz/modules/type/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSettingCreateRequest struct {
	common.SQLModel
	UserId             uuid.UUID   `json:"-" gorm:"column:user_id;"`
	Quantity           int         `json:"quantity" gorm:"column:quantity;"`
	TypeSettingIds     []uuid.UUID `json:"type_setting_ids" gorm:"-"`
	DifficultyId       uuid.UUID   `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingIds []uuid.UUID `json:"category_setting_ids" gorm:"-"`
}

type UserSettingReponse struct {
	common.SQLModel
	UserId     uuid.UUID                   `json:"-" gorm:"column:user_id;"`
	Quantity   int                         `json:"quantity" gorm:"column:quantity;"`
	Types      []typeentity.Type           `json:"types" gorm:"preload:false;"`
	Difficulty difficultyentity.Difficulty `json:"difficulty" gorm:"preload:false;"`
	Categories []categoryentity.Category   `json:"categories" gorm:"preload:false;"`
}

func (usc *UserSettingCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	usc.Id = uuid.UUID(id)

	return err
}

type UserSettingCreate struct {
	common.SQLModel
	UserId            uuid.UUID `json:"-" gorm:"column:user_id;"`
	Quantity          int       `json:"quantity" gorm:"column:quantity;"`
	TypeSettingId     uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
	DifficultyId      uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingId uuid.UUID `json:"category_setting_id" gorm:"column:category_setting_id;"`
}

func (UserSettingCreate) TableName() string {
	return UserSetting{}.TableName()
}

func (usc *UserSettingCreate) Prepare(userId uuid.UUID) {
	usc.UserId = userId
}

type UserSettingUpdate struct {
	common.SQLModel
	UserId            *uuid.UUID `json:"-" gorm:"column:user_id;"`
	Quantity          *int       `json:"quantity" gorm:"column:quantity;"`
	TypeSettingId     *uuid.UUID `json:"type_setting_id" gorm:"column:type_setting_id;"`
	DifficultyId      *uuid.UUID `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategorySettingId *uuid.UUID `json:"category_setting_id" gorm:"column:category_setting_id;"`
}

func (UserSettingUpdate) TableName() string {
	return UserSetting{}.TableName()
}
