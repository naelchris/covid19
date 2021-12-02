package auth

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"io"
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
	UserID int64 `json:"user_id"`
}

type AuthResponse struct {
	UserToken string `json:"user_token"`
	Username     string `json:"username"`
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}