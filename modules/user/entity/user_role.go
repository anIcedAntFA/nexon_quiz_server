package userentity

import (
	"database/sql/driver"
	"fmt"
)

type UserRole int

const (
	RoleAdmin UserRole = iota + 1
	RoleUser
)

func (role UserRole) String() string {
	switch role {
	case RoleAdmin:
		return "admin"
	default:
		return "user"
	}
}

func (role *UserRole) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	var rl UserRole

	roleValue := string(bytes)

	if roleValue == RoleAdmin.String() {
		rl = RoleAdmin
	} else {
		rl = RoleUser
	}

	*role = rl

	return nil
}

func (role *UserRole) Value() (driver.Value, error) {
	if role == nil {
		return nil, nil
	}

	return role.String(), nil
}

func (role *UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", role.String())), nil
}
