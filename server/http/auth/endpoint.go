package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/naelchris/covid19/Internal/repository/user"
	"github.com/naelchris/covid19/server"
	"log"
	"net/http"
	"strings"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	var (
		ctx = r.Context()
		data loginRequest
		userInfo user.UserInfo
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("[authHandler][LoginHandler][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	userInfo, err = server.AuthUsecase.Authenticate(ctx, data.Username, data.Password)
	if err != nil {
		log.Println("[authHandler][LoginHandler][Authenticate] User Authenticate failed, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    userInfo.Name,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		UserID: userInfo.UserID,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		log.Println("[authHandler][LoginHandler][TokenSigned] failed to sign token, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	log.Println("[authHandler][LoginHandler][Success] Successfully generate token.")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	ToJSON(&AuthResponse{UserToken: signedToken, Username: userInfo.Username}, w)
}

func MiddlewareValidateUserToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		tokenParsed, err := extractToken(r)
		if err != nil {
			server.RenderError(w, http.StatusInternalServerError, err, timeStart)
			return
		}
		log.Println(tokenParsed)
		log.Println("After Extract")

		token, err := jwt.ParseWithClaims(tokenParsed, &UserClaims{},func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				log.Println("error 1")
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				log.Println("error 2")
				return nil, fmt.Errorf("Signing method invalid")
			}
			log.Println("return JWT")
			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			server.RenderError(w, http.StatusBadRequest, err, timeStart)
			return
		}
		log.Println(token)
		log.Println("After JWT Parse")
		claims, ok := token.Claims.(*UserClaims)
		if !ok || !token.Valid {
			server.RenderError(w, http.StatusBadRequest, err, timeStart)
			return
		}
		log.Println("After Claim")

		ctx := context.WithValue(r.Context(), "UserID", claims.UserID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 {
		return "", errors.New("Token not provided or malformed")
	}
	return authHeaderContent[1], nil
}