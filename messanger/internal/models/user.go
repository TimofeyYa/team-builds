package models

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
	UpdateAt *string
	CreateAt string
}

type RegistrationUser struct {
	Name     string
	Email    string
	Password string
}
