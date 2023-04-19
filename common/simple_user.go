package common

type SimpleUser struct {
	SQLModel
	Username string `json:"username" gorm:"column:username;"`
	Role     string `json:"role" gorm:"column:role;type:ENUM('user', 'admin')"`
}

func (s *SimpleUser) TableName() string {
	return "users"
}
