package common

import (
	"log"

	"github.com/google/uuid"
)

const (
	DbTypeUser = 4
)

const CurrentUser = "user"

const TimeLayout = "2006-01-02T15:04:05.999999"

type Requester interface {
	GetUserId() uuid.UUID
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
