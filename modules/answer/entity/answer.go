package answerentity

import (
	"fmt"
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
)

const EntityName = "Answer"

type Answer struct {
	common.SQLModel `json:",inline"`
	QuestionId      uuid.UUID `json:"question_id" gorm:"question_id"`
	Content         string    `json:"content" gorm:"column:content"`
	Correct         int       `json:"correct" gorm:"column:correct"`
}

func (Answer) TableName() string {
	return "answers"
}

type AnswerCreate struct {
	common.SQLModel `json:",inline"`
	QuestionId      uuid.UUID `json:"question_id" gorm:"question_id"`
	Content         string    `json:"content" gorm:"column:content"`
	Correct         int       `json:"correct" gorm:"column:correct"`
}

func (*AnswerCreate) TableName() string {
	return "answers"
}

func (ans *AnswerCreate) Validate() error {
	dataContent := ans.Content

	if strings.TrimSpace(dataContent) == "" {
		return common.ErrorInvalidRequest(ErrorFieldIsEmpty("answer content"))
	}

	return nil
}

type AnswerUpdate struct {
	common.SQLModel `json:",inline"`
	QuestionId      uuid.UUID `json:"question_id" gorm:"question_id"`
	Content         string    `json:"content" gorm:"column:content"`
	Correct         int       `json:"correct" gorm:"column:correct"`
}

func (*AnswerUpdate) TableName() string {
	return "answers"
}

type Answers []Answer

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}
