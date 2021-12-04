package auth

import (
	"context"
	"log"

	"github.com/GerardSoleCa/wordpress-hash-go"
	"github.com/naelchris/covid19/Internal/repository/user"
)

type AuthUsecase struct {
	userDomain user.DomainItf
}

func NewAuthUsecase(userDomain user.DomainItf) *AuthUsecase {
	return &AuthUsecase{
		userDomain: userDomain,
	}
}

func (uc *AuthUsecase) Authenticate(ctx context.Context, email string, password string) (user.UserInfo, error) {
	hash := wphash.HashPassword(password)
	req, err := uc.userDomain.GetUser(ctx, email, hash)
	if err != nil {
		log.Println("[AuthUsecase][Login] fail to get user by email err,", err)
		return user.UserInfo{}, err
	}

	userInfo := user.UserInfo{
		Email:        req.Email,
		Name:         req.Name,
		DateOfBirth:  req.DateOfBirth,
		HealthStatus: req.HealthStatus,
	}

	return userInfo, nil
}
