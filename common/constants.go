package common

import (
	"log"

	"github.com/google/uuid"
)

const (
	DbTypeUser = 4
)

const CurrentUser = "user"

const RootAdminRole = "f50fc90b-486e-47f7-89b0-1df44b247861"
const AdminRole = "c39c2f6a-3ac9-4d6b-a21f-fd1ba94eec38"
const UserRole = "839eadd3-5314-49e5-a16b-0229403b68ed"

const TimeLayout = "2006-01-02T15:04:05.999999"

type Requester interface {
	GetUserId() uuid.UUID
	GetEmail() string
	GetRoleId() uuid.UUID
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
