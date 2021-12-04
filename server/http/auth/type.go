package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserClaims struct {
	jwt.StandardClaims
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	HealthStatus string    `json:"health_status"`
}

type AuthResponse struct {
	UserToken    string    `json:"user_token"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	HealthStatus string    `json:"health_status"`
}
