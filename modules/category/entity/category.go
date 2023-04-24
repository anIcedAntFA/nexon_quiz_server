package categoryentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type Category struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) GetCategoryId() uuid.UUID {
	return c.Id
}
