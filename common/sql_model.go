package common

import (
	"time"

	"github.com/google/uuid"
)

type SQLModel struct {
	Id        uuid.UUID  `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time `json:"createdAt,omitempty" gorm:"column:createdAt;"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt;"`
}

type MultipleIDs struct {
	Ids []uuid.UUID `json:"ids,string" gorm:"-"`
}
