package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"nexon_quiz/components/tokenprovider"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.RegisteredClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate jwt
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		},
	})
	//convert key to []bytes
	myToken, err := t.SignedString([]byte(j.secret))

	if err != nil {
		return nil, err
	}

	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		return nil, tokenprovider.ErrorInvalidToken
	}

	// validate token
	if !res.Valid {

		return nil, tokenprovider.ErrorInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)

	if !ok {
		return nil, tokenprovider.ErrorInvalidToken
	}

	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
