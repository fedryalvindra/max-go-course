package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "tungtungsahur"

func GenerateToken(email string, userId int64) (string, error) {
	// Generate new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// best practice not to use password for security reason
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Check token if it SigningMethodHS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return errors.New("could not parsed token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return errors.New("invalid token")
	}

	// Claim email and userId from token
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)

	// if !ok {
	// 	return errors.New("invalid token claims")
	// }

	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
