package structures

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	UserType string `json:"user_type"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}


