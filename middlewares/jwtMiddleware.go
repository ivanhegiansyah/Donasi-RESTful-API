package middlewares

import (
	"finalproject-BE/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userId"] = userId
	claims["name"] = name
	//optional
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //token exp for 1 hours
	//claims["authorized"] = true

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

