package difficultyentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type Difficulty struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (Difficulty) TableName() string {
	return "difficulties"
}

func (d *Difficulty) GetDifficultyId() uuid.UUID {
	return d.Id
}
