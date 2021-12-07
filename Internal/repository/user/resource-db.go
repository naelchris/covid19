package user

import (
	"context"
	"database/sql"
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
		data.Lat,
		data.Lng,
		data.VaccineType,
		data.HealthStatus,
		time.Now(),
	).Scan(&resp.ID, &resp.Name, &resp.Email, &resp.Password, &resp.DateOfBirth, &resp.Lat, &resp.Lng, &resp.VaccineType, &resp.HealthStatus, &resp.CreatedAt)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][addUser] problem query to db err", err.Error())
		return
	}

	return
}

func (s storage) GetUser(ctx context.Context, email string) (resp User, err error) {
	log.Println("[ClassRepository][ResourceDB][getUser] Data Class,", email)

	vaccineCertificate1 := sql.NullString{}
	vaccineCertificate2 := sql.NullString{}
	healthStatus := sql.NullString{}
	vaccineType := sql.NullString{}

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
		&resp.Lat,
		&resp.Lng,
		&vaccineType,
		&vaccineCertificate1,
		&vaccineCertificate2,
		&resp.Password,
		&healthStatus,
		&resp.CreatedAt,
		&resp.UpdatedAt,
	)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][getUser] problem query to db err", err.Error())
		return
	}

	resp.VaccineCertificate1 = vaccineCertificate1.String
	resp.VaccineCertificate2 = vaccineCertificate2.String
	resp.HealthStatus = healthStatus.String
	resp.VaccineType = vaccineType.String

	return
}

func (s storage) ValidateLogin(ctx context.Context, email string, password string) (resp User, err error) {
	log.Println("[ClassRepository][ResourceDB][ValidateLogin] Data Class,", email, password)

	//prepare
	qr, err := s.CasesDB.Prepare(validateLoginQuery)
	if err != nil {
		log.Fatalln("[ClassRepository][ResourceDB][ValidateLogin] prepare failed err, ", err.Error())
		return
	}

	err = qr.QueryRow(
		email,
		password,
	).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Email,
		&resp.DateOfBirth,
		&resp.Lat,
		&resp.Lng,
		&resp.VaccineType,
		&resp.Password,
		&resp.HealthStatus,
	)
	if err != nil {
		log.Println("[ClassRepository][ResourceDB][ValidateLogin] problem query to db err", err.Error())
		return
	}

	return
}

func (s storage) UpdateUser(ctx context.Context, userInfo UserInfo) (resp UserInfo, err error) {
	log.Println(fmt.Sprintf("[UpdateUser domain][UpdateUser][UserInfo: %+v]", userInfo))

	vaccineCertificate1 := sql.NullString{}
	vaccineCertificate2 := sql.NullString{}
	healthStatus := sql.NullString{}
	vaccineType := sql.NullString{}
	lat := sql.NullFloat64{}
	lng := sql.NullFloat64{}

	//construct query
	sql, args := buildUpdateUserQuery(userInfo)
	sql += " RETURNING email, name, dateofbirth, lat, lng, vaccinecertificate1, vaccinecertificate2, healthstatus, vaccinetype"

	log.Println(sql)

	err = s.CasesDB.QueryRow(sqlx.Rebind(2, sql), args...).Scan(
		&resp.Email,
		&resp.Name,
		&resp.DateOfBirth,
		&lat,
		&lng,
		&vaccineCertificate1,
		&vaccineCertificate2,
		&healthStatus,
		&vaccineType,
	)
	if err != nil {
		log.Println("[UpdateUser domain][UpdateUser] Query err, ", err)
		return resp, nil
	}

	resp.VaccineCertificate1 = vaccineCertificate1.String
	resp.VaccineCertificate2 = vaccineCertificate2.String
	resp.HealthStatus = healthStatus.String
	resp.VaccineType = vaccineType.String
	resp.Lat = lat.Float64
	resp.Lng = lng.Float64

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

	updateBuilder.SetMore(updateBuilder.Assign("lat", userInfo.Lat))

	updateBuilder.SetMore(updateBuilder.Assign("lng", userInfo.Lng))

	updateBuilder.SetMore(updateBuilder.Assign("dateofbirth", userInfo.DateOfBirth))

	updateBuilder.SetMore(updateBuilder.Assign("healthstatus", userInfo.HealthStatus))

	updateBuilder.SetMore("updatedat = CURRENT_TIMESTAMP")

	updateBuilder.Where(
		updateBuilder.Equal("email", userInfo.Email),
	)

	sql, args := updateBuilder.Build()

	return sql, args
}
