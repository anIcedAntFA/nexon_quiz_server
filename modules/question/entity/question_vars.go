package questionentity

import (
	"nexon_quiz/common"
	answerentity "nexon_quiz/modules/answer/entity"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QuestionPagingResult struct {
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PageSize     int `json:"page_size"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

func (QuestionPagingResult) TableName() string {
	return Question{}.TableName()
}

func (qc *QuestionCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewUUID()
	qc.Id = uuid.UUID(id)

	for _, answer := range *qc.Answers {
		answer.QuestionId = qc.Id

		id, err := uuid.NewRandom()
		answer.Id = uuid.UUID(id)

		return err
	}

	return err
}

type QuestionCreate struct {
	common.SQLModel
	OwnerId      uuid.UUID                   `json:"-" gorm:"column:owner_id;"`
	Content      string                      `json:"content" gorm:"column:content;"`
	TypeId       uuid.UUID                   `json:"type_id" gorm:"column:type_id;"`
	DifficultyId uuid.UUID                   `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategoryId   uuid.UUID                   `json:"category_id" gorm:"column:category_id;"`
	PlusScore    int                         `json:"plus_score" gorm:"column:plus_score;"`
	MinusScore   int                         `json:"minus_score" gorm:"column:minus_score;"`
	Time         int                         `json:"time" gorm:"column:time;"`
	Answers      *answerentity.AnswersCreate `json:"answers" gorm:"foreignKey:QuestionId"`
}

func (QuestionCreate) TableName() string {
	return Question{}.TableName()
}

func (qc *QuestionCreate) Prepare(
	ownerId uuid.UUID,
	plusCore int,
	minusScore int,
	timePerQuestion int,
) {
	qc.OwnerId = ownerId
	qc.PlusScore = 5
	qc.MinusScore = 5
	qc.Time = 40
}

func (qc *QuestionCreate) Validate() error {
	questionCreateMap := map[string]string{
		"content":       qc.Content,
		"type_id":       qc.TypeId.String(),
		"difficulty_id": qc.CategoryId.String(),
		"category_id":   qc.CategoryId.String(),
	}

	for _, question := range questionCreateMap {
		question = strings.TrimSpace(question)

		if err := checkEmptyField(question); err != nil {
			return err
		}
	}

	return nil
}

type QuestionsCreate = []*QuestionCreate

type QuestionUpdate struct {
	common.SQLModel
	Content      *string `json:"content" gorm:"column:content;"`
	TypeId       *string `json:"type_id" gorm:"column:type_id;"`
	DifficultyId *string `json:"difficulty_id" gorm:"column:difficulty_id;"`
	CategoryId   *string `json:"category_id" gorm:"column:category_id;"`
	PlusScore    *int    `json:"plus_score" gorm:"column:plus_score;"`
	MinusScore   *int    `json:"minus_score" gorm:"column:minus_score;"`
	Time         *int    `json:"time" gorm:"column:time;"`
	Status       *int    `json:"status" gorm:"column:status;default:1;"`
}

func (QuestionUpdate) TableName() string {
	return Question{}.TableName()
}

func (qu *QuestionUpdate) Validate() error {
	questionUpdateMap := map[string]*string{
		"content":       qu.Content,
		"type_id":       qu.TypeId,
		"difficulty_id": qu.CategoryId,
		"category_id":   qu.CategoryId,
	}

	for _, question := range questionUpdateMap {
		str := strings.TrimSpace(*question)

		if err := checkEmptyField(str); err != nil {
			return err
		}

		question = &str
	}

	return nil
}
