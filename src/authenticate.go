package main

import (
	"log"
	"net/http"
	"strings"
)

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := strings.TrimSpace(req.Header.Get("authorization"))
		if authorizationHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) != 2 {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		isTokenValid, err := isJWTTokenValid(bearerToken[1])
		if err != nil {
			log.Printf("isJWTTokenValid error: %s", err)
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		if !isTokenValid {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		next(w, req)
	})
}
