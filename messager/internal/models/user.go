package models

import (
	"time"
)

type User struct {
	Id       int        `db:"id" json:"id"`
	Name     string     `db:"name" json:"name"`
	Email    string     `db:"email" json:"email"`
	Password string     `db:"password_hash" json:"password_hash"`
	UpdateAt *time.Time `db:"update_at" json:"update_at"`
	CreateAt time.Time  `db:"create_at" json:"create_at"`
}

type RegistrationUser struct {
	Name     string
	Email    string
	Password string
}
