package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64        `json:"user_ID", db:"id"`
	Name         string       `json:"name", db:"name"`
	Email        string       `json:"email", db:"email"`
	Password     string       `json:"password", db:"password"`
	DateOfBirth  time.Time    `json:"date_of_birth", db:"dateofbirth"`
	VaccineType  string       `json:"vaccine_type", db:"vaccinetype"`
	HealthStatus string       `json:"health_status", db:"healthstatus"`
	CreatedAt    time.Time    `json:"created_at", db:"createdat"`
	UpdatedAt    sql.NullTime `json:"updated_at", db:"updatedat"`
}

type storage struct {
	CasesDB *sql.DB
}

type UserInfo struct {
	Email               string    `json:"email"`
	Name                string    `json:"name"`
	DateOfBirth         time.Time `json:"date_of_birth"`
	VaccineCertificate1 string    `json:"vaccine_certificate_1"`
	VaccineCertificate2 string    `json:"vaccine_certificate_2"`
	HealthStatus        string    `json:"heatlh_status"`
	VaccineType         string    `json:"vaccine_type"`
}

type UserResponse struct {
	ID int64
}
