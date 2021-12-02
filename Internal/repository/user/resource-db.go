package user

import (
	"context"
	"fmt"
	"log"
	"time"
)

func (s storage) AddUser(ctx context.Context, data User) (resp User, err error) {
	log.Println("[ClassRepository][ResourceDB][addUser] Data Class,", data)

	//prepare
	qr, err := s.CasesDB.Prepare(addUserQuery)
	if err != nil {
		log.Fatalln("[ClassRepository][ResourceDB][addUser] prepare failed err, ", err.Error())
		return resp, err
	}

	err = qr.QueryRow(
		data.Name,
		data.Email,
		data.Password,
		data.DateOfBirth,
		data.VaccineType,
		time.Now(),
	).Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.DateOfBirth, &resp.VaccineType, &resp.CreatedAt, &resp.UpdatedAt)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][addUser] problem query to db err", err.Error())
		return
	}

	return
}

func (s storage) GetUser(ctx context.Context, email string, password string) (resp User, err error) {
	log.Println("[ClassRepository][ResourceDB][getUser] Data Class,", email, password)

	//prepare
	qr, err := s.CasesDB.Prepare(getUserQuery)
	if err != nil {
		log.Fatalln("[ClassRepository][ResourceDB][getUser] prepare failed err, ", err.Error())
		return resp, err
	}

	fmt.Println(qr)
	fmt.Println(email)
	fmt.Println(password)

	err = qr.QueryRow(
		email,
		password,
	).Scan(&resp.ID, &resp.Name, &resp.Email, &resp.DateOfBirth, &resp.VaccineType)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][getUser] problem query to db err", err.Error())
		return
	}

	return
}
