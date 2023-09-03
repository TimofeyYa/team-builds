package service

import (
	"teamBuild/messages/internal/models"
	httpParcer "teamBuilds/libs/http_parcer"
)

type Auth interface {
	LoginUser(models.Credentials) (string, *httpParcer.ErrorHTTP)
	CreateUser(models.RegistrationUser) (*models.User, *httpParcer.ErrorHTTP)
}

type User interface {
}

type Service struct {
	Auth
	User
}

func NewService() *Service {
	return &Service{}
}
