package userroleentity

import (
	"errors"
	"nexon_quiz/common"

	"github.com/google/uuid"
)

type UserRoleContent int

const (
	RoleRootAdmin UserRoleContent = iota
	RoleAdmin
	RoleUser
)

var allUserRoles = [3]string{"root_admin", "admin", "user"}

func (urc *UserRoleContent) String() string {
	return allUserRoles[*urc]
}

func parseStringToUserRoleContent(str string) (UserRoleContent, error) {
	for i := range allUserRoles {
		if allUserRoles[i] == str {
			return UserRoleContent(i), nil
		}
	}

	return UserRoleContent(2), errors.New("invalid user role string")
}

type UserRole struct {
	common.SQLModel
	Content UserRoleContent `json:"content" gorm:"column:content;"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

func (ur *UserRole) GetUserRoleId() uuid.UUID {
	return ur.Id
}

func (ur *UserRole) GetUserRole() int {
	return int(ur.Content)
}
