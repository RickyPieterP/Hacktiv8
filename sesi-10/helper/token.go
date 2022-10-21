package helper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email   string
	Role    string
	Expired time.Time
}

var secret = "IniSecret"

func GenerateToken(payload *Token) (string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tok, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tok, nil
}

func ValidateToken(tokString string) (*Token, error) {
	errResponse := fmt.Errorf(("need signin"))
	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok && !tok.Valid {
		return nil, errResponse
	}

	payload, ok := claims["payload"].(Token)
	if !ok && !tok.Valid {
		return nil, errResponse
	}
	return &payload, nil
}
