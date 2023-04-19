package answerentity

import (
	"fmt"
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const EntityName = "Answer"

type Answer struct {
	common.SQLModel
	QuestionId uuid.UUID `json:"-" gorm:"column:question_id"`
	Content    string    `json:"content" gorm:"column:content"`
	Correct    bool      `json:"correct" gorm:"column:correct"`
	IsDeleted  bool      `json:"-" gorm:"column:is_deleted;"`
}

func (Answer) TableName() string {
	return "answers"
}

type AnswerCreate struct {
	common.SQLModel
	QuestionId uuid.UUID `json:"question_id" gorm:"column:question_id"`
	Content    string    `json:"content" gorm:"column:content"`
	Correct    bool      `json:"correct" gorm:"column:correct"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:is_deleted;"`
}

func (AnswerCreate) TableName() string {
	return Answer{}.TableName()
}

func (ac *AnswerCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	ac.Id = uuid.UUID(id)

	return err
}

type AnswersCreate = []AnswerCreate

func (ac *AnswerCreate) Validate() error {
	dataContent := ac.Content

	if strings.TrimSpace(dataContent) == "" {
		return common.ErrorInvalidRequest(ErrorFieldIsEmpty("answer content"))
	}

	return nil
}

type AnswerUpdate struct {
	common.SQLModel
	QuestionId uuid.UUID `json:"question_id" gorm:"column:question_id"`
	Content    string    `json:"content" gorm:"column:content"`
	Correct    bool      `json:"correct" gorm:"column:correct"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:is_deleted;"`
}

func (*AnswerUpdate) TableName() string {
	return "answers"
}

type Answers []Answer

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
