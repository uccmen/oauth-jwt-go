package main

import (
	"fmt"

	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
)

type CustomJWTClaims struct {
	Sub   string
	Name  string
	Email string
}

func createJWTTokenWithOAuthToken(oauthToken oauth2.Token, claims CustomJWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: oauthToken.Expiry.UnixNano(),
		Subject:   claims.Sub,
	})
	tokenString, error := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if error != nil {
		log.Printf("token signing: %s", error)
		return "", error
	}

	return tokenString, nil
}

func isJWTTokenValid(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		log.Printf("jwt parse: %s", err)
		return false, err
	}

	return token.Valid, nil
}
