package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const SecretKey = "secret"

func GenerateJwt(issuer string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	return claims.SignedString([]byte(SecretKey))
}

func Parsejwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", jwt.ErrSignatureInvalid
	}
	claims := token.Claims.(*jwt.StandardClaims)
	return claims.Issuer, nil
}
