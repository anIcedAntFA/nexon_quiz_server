package userroleentity

import (
	"nexon_quiz/common"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (ur *UserRoleCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	ur.Id = uuid.UUID(id)

	return err
}

type UserRoleCreate struct {
	common.SQLModel
	Content UserRoleContent `json:"content" gorm:"column:content;"`
}

func (UserRoleCreate) TableName() string {
	return UserRole{}.TableName()
}

func (urc *UserRoleCreate) Prepare(deleted_at *time.Time) {
	urc.DeletedAt = nil
}

func (urc *UserRoleCreate) Validate() error {
	if err := checkValidContent(urc.Content); err != nil {
		return err
	}

	return nil
}

type UserRoleUpdate struct {
	common.SQLModel
	Content *UserRoleContent `json:"content" gorm:"column:content;"`
}

func (UserRoleUpdate) TableName() string {
	return UserRole{}.TableName()
}

func (urc *UserRoleUpdate) Validate() error {
	if content := urc.Content; content != nil {
		if err := checkValidContent(*content); err != nil {
			return err
		}

		urc.Content = content
	}

	return nil
}
