package userentity

import (
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type User struct {
	common.SQLModel
	Email    string    `json:"email" gorm:"column:email;"`
	Password string    `json:"-" gorm:"column:password;"`
	Salt     string    `json:"-" gorm:"column:salt;"`
	Username string    `json:"username" gorm:"column:username;"`
	RoleId   uuid.UUID `json:"role_id" gorm:"column:role_id;"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUserId() uuid.UUID {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() uuid.UUID {
	return u.RoleId
}
