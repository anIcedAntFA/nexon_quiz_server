package categoryentity

import (
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryPagingResult struct {
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	NextPage     int `json:"next_page"`
	PageSize     int `json:"page_size"`
	TotalItems   int `json:"total_items"`
	TotalPages   int `json:"total_pages"`
}

func (CategoryPagingResult) TableName() string {
	return Category{}.TableName()
}

func (urc *CategoryCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	urc.Id = uuid.UUID(id)

	return err
}

type CategoryCreate struct {
	common.SQLModel
	Content string `json:"content" gorm:"column:content;"`
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (urc *CategoryCreate) Validate() error {
	content := urc.Content
	content = strings.TrimSpace(content)

	if err := checkEmptyContent(content); err != nil {
		return err
	}

	return nil
}

type CategoriesCreate = []*CategoryCreate

type CategoryUpdate struct {
	common.SQLModel
	Content *string `json:"content" gorm:"column:content;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

func (urc *CategoryUpdate) Validate() error {
	if content := urc.Content; content != nil {
		contentStr := strings.TrimSpace(*content)

		if err := checkEmptyContent(contentStr); err != nil {
			return err
		}

		urc.Content = &contentStr
	}

	return nil
}
