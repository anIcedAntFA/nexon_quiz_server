package typeentity

import (
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypePagingResult struct {
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PageSize     int `json:"page_size"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

func (TypePagingResult) TableName() string {
	return Type{}.TableName()
}

func (tc *TypeCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	tc.Id = uuid.UUID(id)

	return err
}

type TypeCreate struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (TypeCreate) TableName() string {
	return Type{}.TableName()
}

func (urc *TypeCreate) Prepare() {
	urc.DeletedAt = nil
}

func (urc *TypeCreate) Validate() error {
	content := urc.Content
	content = strings.TrimSpace(content)

	if err := checkEmptyContent(content); err != nil {
		return err
	}

	return nil
}

type TypeUpdate struct {
	common.SQLModel
	Content *string `json:"content" gorm:"column:content;"`
}

func (TypeUpdate) TableName() string {
	return Type{}.TableName()
}

func (urc *TypeUpdate) Validate() error {
	if content := urc.Content; content != nil {
		contentStr := strings.TrimSpace(*content)

		if err := checkEmptyContent(contentStr); err != nil {
			return err
		}

		urc.Content = &contentStr
	}

	return nil
}
