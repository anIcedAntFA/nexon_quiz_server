package userroleentity

import (
	"nexon_quiz/common"
)

type UserRoleContent int

const (
	RoleRootAdmin UserRoleContent = iota
	RoleAdmin
	RoleUser
)

var allUserRoles = [3]UserRoleContent{RoleRootAdmin, RoleAdmin, RoleUser}

type UserRole struct {
	common.SQLModel
	Content UserRoleContent `json:"content" gorm:"column:content;"`
}

func (UserRole) TableName() string {
	return "user_roles"
}
