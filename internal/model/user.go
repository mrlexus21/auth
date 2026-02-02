package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64
	Info        UserInfo
	Credentials UserCredentials
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type UserInfo struct {
	Name  string
	Email string
	Role  int32
}

type UserCredentials struct {
	Password        string
	PasswordConfirm string
}

type CreateUser struct {
	Info        UserInfo
	Credentials UserCredentials
}

type UpdateUserInfo struct {
	ID    int64
	Name  *string
	Email *string
}
