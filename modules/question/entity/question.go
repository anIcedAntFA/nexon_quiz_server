package questionentity

import (
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"

	"github.com/google/uuid"
)

const EntityName = "Question"

type Question struct {
	common.SQLModel
	OwnerId      uuid.UUID             `json:"-" gorm:"column:owner_id;"`
	Content      string                `json:"content" gorm:"column:content;"`
	TypeId       string                `json:"type_id" gorm:"column:type_id;"`
	DifficultyId string                `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategoryId   string                `json:"category_id" gorm:"column:category_id;"`
	PlusScore    int                   `json:"plus_score" gorm:"column:plus_score;"`
	MinusScore   int                   `json:"minus_score" gorm:"column:minus_score;"`
	Time         int                   `json:"time" gorm:"column:time;"`
	Answers      *answerentity.Answers `json:"answers" gorm:"preload:false;"`
}

func (Question) TableName() string {
	return "questions"
}

func (q *Question) GetQuestionId() uuid.UUID {
	return q.Id
}
