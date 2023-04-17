package userentity

import (
	"errors"
	"fmt"
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string   `json:"email" gorm:"column:email;"`
	Password        string   `json:"-" gorm:"column:password;"`
	Salt            string   `json:"-" gorm:"column:salt;"`
	UserName        string   `json:"user_name" gorm:"column:user_name;"`
	Role            UserRole `json:"role" gorm:"column:role;type:ENUM('user', 'admin')"`
	Status          int      `json:"status" gorm:"column:status;default:1;"`
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

func (u *User) GetRole() string {
	return u.Role.String()
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string   `json:"email" gorm:"column:email;"`
	Password        string   `json:"password" gorm:"column:password;"`
	UserName        string   `json:"user_name" gorm:"column:user_name;"`
	Salt            string   `json:"-" gorm:"column:salt;"`
	Role            UserRole `json:"-" gorm:"column:role;type:ENUM('user','admin');default:user"`
	Status          int      `json:"status" gorm:"column:status;default:1;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (user *UserCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	user.Id = uuid.UUID(id)

	return err
}

type UserUpdate struct {
	common.SQLModel `json:",inline"`
	Password        *string `json:"password" gorm:"column:password;"`
	UserName        *string `json:"user_name" gorm:"column:user_name;"`
	Salt            string  `json:"-" gorm:"column:salt;"`
	Status          int     `json:"status" gorm:"column:status;default:1;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (u *UserLogin) TableName() string {
	return User{}.TableName()
}

func (data *UserCreate) Validate() error {
	dataNames := map[string]string{
		"user_name": data.UserName,
		"email":     data.Email,
		"password":  data.Password,
	}

	for k, v := range dataNames {
		v = strings.TrimSpace(v)

		if v == "" {
			return ErrorFieldIsEmpty(k)
		}
	}

	return nil
}

func ErrorFieldIsEmpty(field string) error {
	return fmt.Errorf("%s cannot be empty", field)
}

var (
	ErrorEmailOrPasswordInvalid = common.NewCustomError(
		errors.New("email or password invalid"),
		"email or password invalid",
		"ErrorEmailOrPasswordInvalid",
	)
	ErrorEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrorEmailExisted",
	)
	ErrorUserDisabledOrBanned = common.NewCustomError(
		errors.New("user has been disabled or banned"),
		"user has been disabled or banned",
		"ErrorUserDisabledOrBanned",
	)
)
