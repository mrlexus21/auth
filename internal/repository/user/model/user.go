package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64           `db:"id"`
	Info        UserInfo        `db:""`
	Credentials UserCredentials `db:""`
	CreatedAt   time.Time       `db:"created_at"`
	UpdatedAt   sql.NullTime    `db:"updated_at"`
}

type UserInfo struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	Role  int32  `db:"role"`
}

type UserCredentials struct {
	Password        string `db:"password"`
	PasswordConfirm string `db:"password_confirm"`
}

type CreateUser struct {
	Info        UserInfo        `db:""`
	Credentials UserCredentials `db:""`
}

type UpdateUserInfo struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}
