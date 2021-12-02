package auth

import (
	"context"
	"errors"
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
	req, err := uc.userDomain.GetUserByEmail(ctx, email)
	if err != nil {
		log.Println("[AuthUsecase][Login] fail to get user by email err,", err)
		return user.UserInfo{}, err
	}

	hash := wphash.HashPassword(password)
	checked := wphash.CheckPassword(req.Password, hash)

	if !checked {
		err = errors.New("[password not match")
		log.Println("[AuthUsecase][Login] wrong password,", err)
		return user.UserInfo{}, err
	}

	userInfo := user.UserInfo{
		UserID: req.UserID,
		Username: req.Username,
		Email: req.Email,
		Name: req.Name,
	}

	return userInfo, nil
}
