package user

import (
	"context"
	"database/sql"
	"log"
)

type DomainItf interface {
	AddUser(ctx context.Context, data User) (User, error)
	GetUser(ctx context.Context, email string, password string) (User, error)
	UpdateUser(ctx context.Context, userInfo UserInfo) (resp UserInfo, err error)
}

type DBResourceItf interface {
	AddUser(ctx context.Context, data User) (User, error)
	GetUser(ctx context.Context, email string, password string) (User, error)
	UpdateUser(ctx context.Context, userInfo UserInfo) (resp UserInfo, err error)
}

type Domain struct {
	Storage DBResourceItf
}

func InitDomain(db *sql.DB) Domain {
	return Domain{
		Storage: storage{
			CasesDB: db,
		},
	}
}

func (d Domain) AddUser(ctx context.Context, data User) (resp User, err error) {
	resp, err = d.Storage.AddUser(ctx, data)
	if err != nil {
		log.Println("[ClassUsecase][AddUser] problem when querying to database, err :", err)
		return
	}

	return
}

func (d Domain) GetUser(ctx context.Context, email string, password string) (resp User, err error) {
	resp, err = d.Storage.GetUser(ctx, email, password)
	if err != nil {
		log.Println("[ClassUsecase][GetUser] problem when querying to database, err :", err)
		return
	}

	return
}

func (d Domain) UpdateUser(ctx context.Context, userInfo UserInfo) (resp UserInfo, err error) {
	resp, err = d.Storage.UpdateUser(ctx, userInfo)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
