package service

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
)

type Auth interface {
	LoginUser(context.Context, models.Credentials) (string, *httpParcer.ErrorHTTP)
	CreateUser(context.Context, models.RegistrationUser) (*models.User, *httpParcer.ErrorHTTP)
}

type User interface {
}

type Service struct {
	Auth
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
