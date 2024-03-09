package config

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte("98irh89h94hfbhfguvdfs3")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}