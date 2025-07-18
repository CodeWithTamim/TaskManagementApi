package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env, %v", err)
	}
}

func CreateJWTClaim(email string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	tokenString, err := token.SignedString([]byte(secret))

	return tokenString, err
}

func ParseJWTClaim(tokenString string) (string, int64, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return "", 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return "", 0, errors.New("Claim is not valid")
	}

	//extract email
	//
	email, ok := claims["email"]

	if !ok {
		return "", 0, errors.New("Missing email or invalid.")
	}

	exp, ok := claims["exp"]

	if !ok {
		return "", 0, errors.New("Expiry email or invalid.")
	}

	return email.(string), int64(exp.(float64)), nil
}
