package initializers

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("SECRET"))

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}