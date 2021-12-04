package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int64          `json:"user_ID", db:"id"`
	Name         string         `json:"name", db:"name"`
	Email        string         `json:"email", db:"email"`
	Password     string         `json:"password", db:"password"`
	DateOfBirth  time.Time      `json:"date_of_birth", db:"dateofbirth"`
	VaccineType  sql.NullString `json:"vaccine_type", db:"vaccinetype"`
	HealthStatus sql.NullString `json:"health_status", db:"healthstatus"`
	CreatedAt    time.Time      `json:"created_at", db:"createdat"`
	UpdatedAt    sql.NullTime   `json:"updated_at", db:"updatedat"`
}

type storage struct {
	CasesDB *sql.DB
}

type UserInfo struct {
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	HealthStatus string    `json:"heatlh_status"`
}

type UserResponse struct {
	ID int64
}
