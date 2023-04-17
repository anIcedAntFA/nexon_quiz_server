package tokenprovider

import (
	"errors"
	"nexon_quiz/common"
	"time"

	"github.com/google/uuid"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
}

var (
	ErrorNotFound = common.NewCustomError(
		errors.New("token not found"),
		"token not found",
		"ErrorNotFound",
	)

	ErrorEncodingToken = common.NewCustomError(
		errors.New("error encoding token"),
		"error encoding token",
		"ErrorEncodingToken",
	)

	ErrorInvalidToken = common.NewCustomError(
		errors.New("invalid token provided"),
		"invalid token provided",
		"ErrorInvalidToken",
	)
)
