package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func GetAccessToken(userRole string) (string, error) {
	payload := jwt.MapClaims{
		"role": userRole,
		"exp":  time.Now().Add(time.Hour * 72),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	secureKey := os.Getenv("SECURE_KEY")
	accessToken, err := token.SignedString([]byte(secureKey))

	if err != nil {
		log.Printf("Get access token error: %s", err.Error())
		return "", err
	}

	return accessToken, nil
}
