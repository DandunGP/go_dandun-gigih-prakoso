package middleware

import (
	"belajar-go-echo/constant"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId int, email string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constant.SECRET_KEY))
}
