package userentity

import (
	"nexon_quiz/common"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (uc *UserCreate) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewRandom()
	uc.Id = uuid.UUID(id)

	return err
}

type UserCreate struct {
	common.SQLModel
	Email    string    `json:"email" gorm:"column:email;"`
	Password string    `json:"password" gorm:"column:password;"`
	Username string    `json:"username" gorm:"column:username;"`
	Salt     string    `json:"-" gorm:"column:salt;"`
	RoleId   uuid.UUID `json:"-" gorm:"column:role_id;"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

// func (uc *UserCreate) Prepare(userRoleId uuid.UUID, password string) {
// 	uc.RoleId = userRoleId
// 	uc.Password = password
// 	uc.DeletedAt = nil
// }

func (uc *UserCreate) Validate() error {
	userCreateMap := map[string]string{
		"username": uc.Username,
		"email":    uc.Email,
		"password": uc.Password,
	}

	for _, user := range userCreateMap {
		user = strings.TrimSpace(user)

		if err := checkEmptyField(user); err != nil {
			return err
		}
	}

	return nil
}

type UserUpdate struct {
	common.SQLModel
	Password *string   `json:"password" gorm:"column:password;"`
	Salt     string    `json:"-" gorm:"column:salt;"`
	Username *string   `json:"username" gorm:"column:username;"`
	RoleId   uuid.UUID `json:"-" gorm:"column:role_id;"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

func (uu *UserUpdate) Validate() error {
	userUpdateMap := map[string]*string{
		"username": uu.Username,
		"password": uu.Password,
	}

	for _, user := range userUpdateMap {
		str := strings.TrimSpace(*user)

		if err := checkEmptyField(str); err != nil {
			return err
		}

		user = &str

	}

	return nil
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
