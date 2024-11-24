package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const superKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(superKey))
}

func VerifyToken(token string) (int64, error){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("UNEXPECTED SIGNING METHOD")
		}
		
		return []byte(superKey), nil
	})

	if err != nil {
		return 0, errors.New("COULD NOT PARSE TOKEN")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("INVALID TOKEN")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("INVALID TOKEN CLAIMS")
	}

	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}