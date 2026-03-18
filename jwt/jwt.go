package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

func GenToken(claims jwt.MapClaims, hmacSampleSecret string) string {
	tokenString, err := GenTokenE(claims, hmacSampleSecret)
	if err != nil {
		return ""
	}
	return tokenString
}

func GenTokenE(claims jwt.MapClaims, hmacSampleSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(hmacSampleSecret))
}

func ParseToken(tokenString string, hmacSampleSecret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(hmacSampleSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
