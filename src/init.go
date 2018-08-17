package main

import (
	"log"

	"crypto/rand"
	"encoding/hex"

	"os"
)

func uuid() string {
	b := make([]byte, 12)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalln("unable to create uuid")
	}
	return hex.EncodeToString(b)
}

func init() {
	if os.Getenv("PORT") == "" {
		log.Fatalln("PORT not set")
	}
	if os.Getenv("FB_CLIENT_ID") == "" {
		log.Fatalln("FB_CLIENT_ID not set")
	}
	if os.Getenv("FB_CLIENT_SECRET") == "" {
		log.Fatalln("FB_CLIENT_SECRET not set")
	}
	if os.Getenv("APP_CALLBACK_REDIRECT_URL") == "" {
		log.Fatalln("APP_CALLBACK_REDIRECT_URL not set")
	}
	if os.Getenv("JWT_SECRET_KEY") == "" {
		log.Fatalln("JWT_SECRET_KEY not set")
	}
}
