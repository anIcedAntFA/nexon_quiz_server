package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

type SimpleSQLModel struct {
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

func NewSQLModel(newId uuid.UUID) SQLModel {
	now := time.Now().UTC()

	return SQLModel{
		Id:        newId,
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}

type MultipleIDs struct {
	Ids []uuid.UUID `json:"ids" gorm:"-"`
}
