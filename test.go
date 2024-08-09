package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const tokenTTL = time.Hour * 720

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID   int64
	Email    string
	UserType string
}

func GenerateJWT(jwtKey string, userID int64, email string, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTTL)},
		},
		userID,
		email,
		userType,
	})

	return token.SignedString([]byte(jwtKey))
}

func GetTokenClaims(inputToken string, key string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(inputToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

func main() {
	jwtKey := "123"

	token, err := GenerateJWT(jwtKey, 10, "test@mail.ru", "client")
	if err != nil {
		return
	}

	claims, err := GetTokenClaims(token, "huita")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(token, claims)
}
