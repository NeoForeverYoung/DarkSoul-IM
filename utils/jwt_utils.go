package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var jwtKey = []byte("your_secret_key")

type Claims struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateToken(userId uint) (string, error) {
	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ParseToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if token.Valid {
		return claims, nil
	}
	return nil, err
} 