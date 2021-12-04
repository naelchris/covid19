package auth

import (
	"context"
	"errors"
	"github.com/GerardSoleCa/wordpress-hash-go"
	"github.com/naelchris/covid19/Internal/repository/user"
	"log"
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
		log.Println("[AuthUsecase][Login] fail to get user err,", err)
		return user.UserInfo{}, err
	}

	userInfo := user.UserInfo{
		Email:        req.Email,
		Name:         req.Name,
		DateOfBirth:  req.DateOfBirth,
		HealthStatus: req.HealthStatus.String,
	}

	return userInfo, nil
}

//TODO: Need to discuss wheter we need to save the token or not in DB
func (uc *AuthUsecase) Register(ctx context.Context, data user.User) (user.UserInfo, error) {

	err := RegisterValidation(data)
	if err != nil {
		log.Println("[AuthUsecase][Register] fail to validate user data err,", err)
		return user.UserInfo{}, err
	}

	data.Password = wphash.HashPassword(data.Password)
	log.Println(data.Password)
	_, err = uc.userDomain.AddUser(ctx, data)
	if err != nil {
		log.Println("[AuthUsecase][Register] fail to add user err,", err)
		return user.UserInfo{}, err
	}

	userInfo := user.UserInfo{
		Email:        data.Email,
		Name:         data.Name,
		DateOfBirth:  data.DateOfBirth,
		HealthStatus: data.HealthStatus.String,
	}

	return userInfo, nil
}

//TODO: Need to add more validation
func RegisterValidation(data user.User) error {
	if data.Name == "" || data.Password == "" || data.Email == "" || data.DateOfBirth.IsZero() || data.CreatedAt.IsZero() {
		return errors.New("There's field that not filled")
	}

	return nil
}
