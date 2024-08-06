package common

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID   int64
	Email    string
	UserType string
}

func GenerateJWT(jwtKey string, userID int64, email string, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.RegisteredClaims{
			IssuedAt: &jwt.NumericDate{Time: time.Now()},
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
