package common

import "github.com/google/uuid"

type SimpleUser struct {
	SQLModel
	Username string    `json:"username" gorm:"column:username;"`
	RoleId   uuid.UUID `json:"role_id" gorm:"column:role_id;"`
}

func (s *SimpleUser) TableName() string {
	return "users"
}
