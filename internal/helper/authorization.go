package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateAccessToken(id int, name string, email string) (string, error) {
	expiration := time.Now().Add(10 * time.Minute)
	claims := jwt.MapClaims{
		"id":    id,
		"name":  name,
		"email": email,
		"exp":   expiration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
