package difficultyentity

import (
	"nexon_quiz/common"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DifficultyPagingResult struct {
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PageSize     int `json:"page_size"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

func (DifficultyPagingResult) TableName() string {
	return Difficulty{}.TableName()
}

func (dc *DifficultyCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	dc.Id = uuid.UUID(id)

	return err
}

type DifficultyCreate struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (DifficultyCreate) TableName() string {
	return Difficulty{}.TableName()
}

func (dc *DifficultyCreate) Prepare(deleted_at *time.Time) {
	dc.DeletedAt = nil
}

func (dc *DifficultyCreate) Validate() error {
	content := dc.Content
	content = strings.TrimSpace(content)

	if err := checkEmptyContent(content); err != nil {
		return err
	}

	return nil
}

type DifficultyUpdate struct {
	common.SQLModel
	Content *string `json:"content" gorm:"column:content;"`
}

func (DifficultyUpdate) TableName() string {
	return Difficulty{}.TableName()
}

func (du *DifficultyUpdate) Validate() error {
	if content := du.Content; content != nil {
		contentStr := strings.TrimSpace(*content)

		if err := checkEmptyContent(contentStr); err != nil {
			return err
		}

		du.Content = &contentStr
	}

	return nil
}
