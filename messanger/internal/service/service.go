package service

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
)

type Auth interface {
	LoginUser(context.Context, models.Credentials) (string, string, *httpParcer.ErrorHTTP)
	CreateUser(context.Context, models.RegistrationUser) (*models.User, *httpParcer.ErrorHTTP)
	Authorization(c context.Context, tokens *models.TokenPair) (*models.TokenPair, *httpParcer.ErrorHTTP)
	ValidateToken(context.Context, string) *httpParcer.ErrorHTTP
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
