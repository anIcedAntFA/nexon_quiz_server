package typeentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type Type struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (Type) TableName() string {
	return "types"
}

func (t *Type) GetCategoryId() uuid.UUID {
	return t.Id
}
