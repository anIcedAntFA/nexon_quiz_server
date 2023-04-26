package answerentity

import (
	"database/sql/driver"
	"errors"
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
	Correct    BitBool   `json:"-" gorm:"column:correct"`
}

func (Answer) TableName() string {
	return "answers"
}

func (ac *AnswerCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	ac.Id = uuid.UUID(id)

	return err
}

type Answers = []Answer

type AnswerCreate struct {
	common.SQLModel
	QuestionId uuid.UUID `json:"question_id" gorm:"column:question_id"`
	Content    string    `json:"content" gorm:"column:content"`
	Correct    BitBool   `json:"correct" gorm:"column:correct"`
}

func (AnswerCreate) TableName() string {
	return Answer{}.TableName()
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
	Content    *string   `json:"content" gorm:"column:content"`
	Correct    *BitBool  `json:"correct" gorm:"column:correct"`
}

func (*AnswerUpdate) TableName() string {
	return "answers"
}

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}

// BitBool is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type BitBool bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (bb BitBool) Value() (driver.Value, error) {
	if bb {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (bb *BitBool) Scan(src interface{}) error {
	v, ok := src.([]byte)

	if !ok {
		return errors.New("bad []byte type assertion")
	}

	*bb = v[0] == 1

	return nil
}
