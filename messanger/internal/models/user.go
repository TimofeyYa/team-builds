package models

import (
	"time"
)

type User struct {
	Id       int64      `db:"id"`
	Name     string     `db:"name"`
	Email    string     `db:"email"`
	Password string     `db:"password_hash"`
	UpdateAt *time.Time `db:"update_at"`
	CreateAt time.Time  `db:"create_at"`
}

type RegistrationUser struct {
	Name     string
	Email    string
	Password string
}
