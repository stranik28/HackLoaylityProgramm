package helper

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/stranik28/HackLoaylityProgramm/internal/logger"
	"os"
	"time"
)

var SecretKey = GetEnv("SECRET_KEY", "NOT FOUND")

type SignedDetails struct {
	Id int
	jwt.StandardClaims
}

func Init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		logger.Log.Fatal("No .env file found")
	}
}

// GetEnv finds an env variable or the given fallback.
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
		logger.Log.Warn("Not found " + key)
	}
	return value
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("Your Token is not Valid anymore")
		msg = err.Error()
		return
	}
	fmt.Println(claims, msg)
	return claims, msg
}
