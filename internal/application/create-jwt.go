package application

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createJWT(expiresAt time.Duration, userId, jwtSecret string) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresAt)),
		Subject:   userId,
	}).SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
