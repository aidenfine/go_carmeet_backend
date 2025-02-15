package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    userId,
		"expiredAt": time.Now().Add(30).Unix(),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "Error signing jwt", err
	}
	return tokenString, err

}
