package service

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaims struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret_key"
	}
	return secret
}

func GenerateToken(userID string) string {
	claims := &authCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(getSecretKey()))
	if err != nil {
		panic(err)
	}

	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}
		return []byte(getSecretKey()), nil
	})
}

func ParseToken(token *jwt.Token) (*authCustomClaims, bool) {
	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}
