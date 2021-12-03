package user

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
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
		data.HealthStatus,
		time.Now(),
	).Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.DateOfBirth, &resp.VaccineType, &resp.HealthStatus, &resp.CreatedAt)
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

	err = qr.QueryRow(
		email,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Email,
		&resp.DateOfBirth,
		&resp.VaccineType,
		&resp.Password,
		&resp.HealthStatus,
	)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][getUser] problem query to db err", err.Error())
		return
	}

	return
}

func (s storage) UpdateUser(ctx context.Context, userInfo UserInfo) (resp UserInfo, err error) {
	log.Println(fmt.Sprintf("[UpdateUser domain][UpdateUser][UserInfo: %+v]", userInfo))

	//construct query
	sql, args := buildUpdateUserQuery(userInfo)
	sql += " RETURNING email, name, dateofbirth, vaccinecertificate1, vaccinecertificate2, healthstatus, vaccinetype"

	log.Println(sql)

	err = s.CasesDB.QueryRow(sqlx.Rebind(2, sql), args...).Scan(
		&resp.Email,
		&resp.Name,
		&resp.DateOfBirth,
		&resp.VaccineCertificate1,
		&resp.VaccineCertificate2,
		&resp.HealthStatus,
		&resp.VaccineType,
	)
	if err != nil {
		log.Println("[UpdateUser domain][UpdateUser] Query err, ", err)
		return resp, nil
	}

	return resp, nil
}

func buildUpdateUserQuery(userInfo UserInfo) (string, []interface{}) {
	//construct query
	updateBuilder := sqlbuilder.NewUpdateBuilder()
	updateBuilder.Update("user_data")

	if userInfo.VaccineCertificate1 != "" {
		updateBuilder.SetMore(
			updateBuilder.Assign("vaccinecertificate1", userInfo.VaccineCertificate1),
		)
	}

	if userInfo.VaccineCertificate2 != "" {
		updateBuilder.SetMore(updateBuilder.Assign("vaccinecertificate2", userInfo.VaccineCertificate2))
	}

	updateBuilder.SetMore(updateBuilder.Assign("name", userInfo.Name))

	updateBuilder.SetMore(updateBuilder.Assign("dateofbirth", userInfo.DateOfBirth))

	updateBuilder.SetMore(updateBuilder.Assign("healthstatus", userInfo.HealthStatus))

	updateBuilder.SetMore("updatedat = CURRENT_TIMESTAMP")

	updateBuilder.Where(
		updateBuilder.Equal("email", userInfo.Email),
	)

	sql, args := updateBuilder.Build()

	return sql, args
}
