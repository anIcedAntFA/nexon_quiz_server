package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;"`
	IsDeleted int        `json:"is_deleted" gorm:"column:is_deleted;"`
	CreatedAt *time.Time `json:"created_at,omitempty" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at;"`
}

type MultipleIDs struct {
	Ids []uuid.UUID `json:"ids" gorm:"-"`
}
