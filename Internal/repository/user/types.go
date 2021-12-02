package user

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64          `json:"user_ID", db:"id"`
	Name        string         `json:"name", db:"name"`
	Email       string         `json:"email", db:"email"`
	Password    string         `json:"password", db:"password"`
	DateOfBirth time.Time      `json:"date_of_birth", db:"dateofbirth"`
	VaccineType sql.NullString `json:"vaccine_type", db:"vaccine_type"`
	CreatedAt   time.Time      `json:"created_at", db:"createdat"`
	UpdatedAt   sql.NullTime   `json:"updated_at", db:"updatedat"`
}

type storage struct {
	CasesDB *sql.DB
}

type UserInfo struct {
	UserID         int64     `json:"user_id"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
}

type UserResponse struct {
	UserID int64
}

