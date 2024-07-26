package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID   uint
	Username string
	jwt.StandardClaims
}

func GenerateToken(id uint, username string) (string, error) {
	expirationTime := time.Now().Add(64 * time.Hour)
	claims := &Claims{
		UserID:   id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(signedToken string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(signedToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorSignatureInvalid) != 0 {
				return nil, err
			}
		}
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

func IsTokenValid(signedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
