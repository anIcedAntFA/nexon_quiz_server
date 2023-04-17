package common

type SimpleUser struct {
	SQLModel
	UserName string `json:"user_name" gorm:"column:user_name;"`
	Role     string `json:"role" gorm:"column:role;type:ENUM('user', 'admin')"`
}

func (s *SimpleUser) TableName() string {
	return "users"
}
