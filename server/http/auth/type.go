package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

const LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour

var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserClaims struct {
	jwt.StandardClaims
	UserID       int64     `json:"user_id"`
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
