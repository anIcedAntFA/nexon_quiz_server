package questionentity

import (
	"fmt"
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
	"strings"

	"github.com/google/uuid"
)

const EntityName = "Question"

type Question struct {
	common.SQLModel `json:",inline"`
	OwnerId         uuid.UUID             `json:"-" gorm:"column:owner_id;"`
	Content         string                `json:"content" gorm:"column:content;"`
	Category        string                `json:"category" gorm:"column:category;"`
	Type            QuestionType          `json:"type" gorm:"column:type;"`
	Difficulty      QuestionDifficulty    `json:"difficulty" gorm:"column:difficulty;"`
	PlusScore       int                   `json:"plus_score" gorm:"column:plus_score"`
	MinusScore      int                   `json:"minus_score" gorm:"column:minus_score"`
	Time            int                   `json:"time" gorm:"column:time"`
	Status          int                   `json:"status" gorm:"column:status;default:1;"`
	Answers         *answerentity.Answers `json:"answers" gorm:"preload:false"`
}

func (Question) TableName() string {
	return "questions"
}

func (q *Question) GetQuestionId() uuid.UUID {
	return q.Id
}

type QuestionCreate struct {
	common.SQLModel `json:",inline"`
	OwnerId         uuid.UUID          `json:"-" gorm:"column:owner_id;"`
	Content         string             `json:"content" gorm:"column:content;"`
	Category        string             `json:"category" gorm:"column:category;"`
	Type            QuestionType       `json:"type" gorm:"column:type;"`
	Difficulty      QuestionDifficulty `json:"difficulty" gorm:"column:difficulty;"`
	PlusScore       int                `json:"plus_score" gorm:"column:plus_score"`
	MinusScore      int                `json:"minus_score" gorm:"column:minus_score"`
	Time            int                `json:"time" gorm:"column:time"`
}

func (QuestionCreate) TableName() string {
	return Question{}.TableName()
}

func (qc *QuestionCreate) Validate() error {
	dataNames := map[string]string{
		"content":  qc.Content,
		"category": qc.Category,
	}

	for k, v := range dataNames {
		v = strings.TrimSpace(v)

		if v == "" {
			return ErrorFieldIsEmpty(k)
		}
	}

	return nil
}

type QuestionUpdate struct {
	common.SQLModel `json:",inline"`
	Content         *string               `json:"content" gorm:"column:content;"`
	Category        *string               `json:"category" gorm:"column:category;"`
	Type            *QuestionType         `json:"type" gorm:"column:type;"`
	Difficulty      *QuestionDifficulty   `json:"difficulty" gorm:"column:difficulty;"`
	PlusScore       *int                  `json:"plus_score" gorm:"column:plus_score"`
	MinusScore      *int                  `json:"minus_score" gorm:"column:minus_score"`
	Time            *int                  `json:"time" gorm:"column:time"`
	Status          *int                  `json:"status" gorm:"column:status;default:1;"`
	Answers         *answerentity.Answers `json:"answers" gorm:"preload:false"`
}

func (QuestionUpdate) TableName() string {
	return Question{}.TableName()
}

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
