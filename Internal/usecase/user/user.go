package covid

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"

	"github.com/naelchris/covid19/Internal/repository/user"
)

type UserUsecase struct {
	userDomain user.DomainItf
}

func NewUserUsecase(userDomain user.DomainItf) *UserUsecase {
	return &UserUsecase{
		userDomain: userDomain,
	}
}

func (uc *UserUsecase) AddUser(ctx context.Context, userData user.User) (resp user.User, err error) {
	sha := sha1.New()
	sha.Write([]byte(userData.Password))

	encryptedPassword := fmt.Sprintf("%x", sha.Sum(nil))

	userData.Password = encryptedPassword

	resp, err = uc.userDomain.AddUser(ctx, userData)
	if err != nil {
		log.Println("[UserUsecase][AddUser] fail to create user err:", err)
		return
	}

	return
}

func (uc *UserUsecase) GetUser(ctx context.Context, email string, password string) (resp user.User, err error) {
	sha := sha1.New()
	sha.Write([]byte(password))

	encryptedPassword := fmt.Sprintf("%x", sha.Sum(nil))

	resp, err = uc.userDomain.GetUser(ctx, email, encryptedPassword)
	if err != nil {
		log.Println("[UserUsecase][GetUser] fail to get user err:", err)
		return
	}

	return
}
