package covid

import (
	"context"
	"crypto/sha1"
	"fmt"
	"log"
	"mime/multipart"
	"reflect"
	"time"

	"github.com/naelchris/covid19/Internal/config"
	"github.com/naelchris/covid19/Internal/repository/firestore"
	"github.com/naelchris/covid19/Internal/repository/user"
)

type UserUsecase struct {
	userDomain user.DomainItf
	fireStore  firestore.DomainItf
	cfg        config.Config
}

func NewUserUsecase(userDomain user.DomainItf, fireStore firestore.DomainItf, cfg config.Config) *UserUsecase {
	return &UserUsecase{
		userDomain: userDomain,
		fireStore:  fireStore,
		cfg:        cfg,
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

//Hit the google cloud storage for photo
func (uc *UserUsecase) UpdateUser(ctx context.Context, userData user.UserInfo, files map[string]multipart.File) (resp user.UserInfo, err error) {

	//set the data
	resp = userData

	//send to gcs
	for k, f := range files {
		fileName := fmt.Sprintf("%+v:%+v", resp.Email, time.Now().Nanosecond())

		url, err := uc.UploadCertificate(ctx, fileName, f)
		if err != nil {
			return resp, err
		}

		if reflect.DeepEqual(k, "certificate_1") {
			resp.VaccineCertificate1 = fileName
		}

		if reflect.DeepEqual(k, "certificate_2") {
			resp.VaccineCertificate2 = fileName
		}

		log.Println(url)
	}

	log.Println(resp)

	//update user data
	resp, err = uc.userDomain.UpdateUser(ctx, resp)
	if err != nil {
		log.Println("[UpdateUser usecase][UpdateUser] err,", err)
		return resp, err
	}

	resp.VaccineCertificate1, err = uc.fireStore.GenerateV4PutObjectSignedURL(ctx, uc.cfg.Server.BucketName, resp.VaccineCertificate1, uc.cfg.Conf)
	if err != nil {
		return resp, err
	}

	resp.VaccineCertificate2, err = uc.fireStore.GenerateV4PutObjectSignedURL(ctx, uc.cfg.Server.BucketName, resp.VaccineCertificate2, uc.cfg.Conf)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (uc *UserUsecase) UploadCertificate(ctx context.Context, fileName string, file multipart.File) (string, error) {
	//send to gcs
	url, err := uc.fireStore.Upload(ctx, fileName, file)
	if err != nil {
		return "", err
	}

	return url, nil
}
