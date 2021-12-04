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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeStart := time.Now()
	var (
		ctx      = r.Context()
		data     loginRequest
		userInfo user.UserInfo
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("[authHandler][LoginHandler][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	userInfo, err = server.AuthUsecase.Authenticate(ctx, data.Email, data.Password)
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
		Name:         userInfo.Name,
		Email:        userInfo.Email,
		DateOfBirth:  userInfo.DateOfBirth,
		HealthStatus: userInfo.HealthStatus,
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

	resp := AuthResponse{
		Name:         userInfo.Name,
		Email:        userInfo.Email,
		DateOfBirth:  userInfo.DateOfBirth,
		HealthStatus: userInfo.HealthStatus,
		UserToken:    signedToken,
	}

	log.Println("[authHandler][LoginHandler][Success] Successfully generate token.")

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	timeStart := time.Now()
	var (
		ctx      = r.Context()
		data     user.User
		userInfo user.UserInfo
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("[authHandler][RegisterHandler][jsonNewDocoder] decode error err, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	userInfo, err = server.AuthUsecase.Register(ctx, data)
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
		Name:         userInfo.Name,
		Email:        userInfo.Email,
		DateOfBirth:  userInfo.DateOfBirth,
		HealthStatus: userInfo.HealthStatus,
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		log.Println("[authHandler][RegisterHandler][TokenSigned] failed to sign token, ", err)
		server.RenderError(w, http.StatusInternalServerError, err, timeStart)
		return
	}

	resp := AuthResponse{
		Name:         userInfo.Name,
		Email:        userInfo.Email,
		DateOfBirth:  userInfo.DateOfBirth,
		HealthStatus: userInfo.HealthStatus,
		UserToken:    signedToken,
	}

	log.Println("[authHandler][RegisterHandler][Success] Successfully generate token.")

	server.RenderResponse(w, http.StatusOK, resp, timeStart)
}

func MiddlewareValidateUserToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		tokenParsed, err := extractToken(r)
		if err != nil {
			server.RenderError(w, http.StatusInternalServerError, err, timeStart)
			return
		}

		token, err := jwt.ParseWithClaims(tokenParsed, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Signing method invalid")
			} else if method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Signing method invalid")
			}
			log.Println("return JWT")
			return JWT_SIGNATURE_KEY, nil
		})
		if err != nil {
			server.RenderError(w, http.StatusBadRequest, err, timeStart)
			return
		}

		claims, ok := token.Claims.(*UserClaims)
		if !ok || !token.Valid {
			err = errors.New("Token Invalid")
			server.RenderError(w, http.StatusBadRequest, err, timeStart)
			return
		}

		data := user.UserInfo{
			Name:         claims.Name,
			Email:        claims.Email,
			DateOfBirth:  claims.DateOfBirth,
			HealthStatus: claims.HealthStatus,
		}

		ctx := context.WithValue(r.Context(), "data", data)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderContent := strings.Split(authHeader, " ")
	if len(authHeaderContent) != 2 || authHeaderContent[0] != "Bearer" {
		return "", errors.New("Token not provided or malformed")
	}
	return authHeaderContent[1], nil
}

func MiddlewareValidateSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		tokenParsed, err := extractToken(r)
		if err != nil && tokenParsed == "" {
			next.ServeHTTP(w, r)
			return
		}

		err = errors.New("User already login")
		server.RenderError(w, http.StatusForbidden, err, timeStart)
		return
	})
}
