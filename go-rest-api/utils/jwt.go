package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "tungtungsahur"

func GenerateToken(email string, userId int64) (string, error) {
	// Generate new token 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// best practice not to use password for security reason
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}