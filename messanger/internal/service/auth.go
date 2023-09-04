package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
)

const solt = "message"

type AuthService struct {
	repo *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) LoginUser(c context.Context, cred models.Credentials) (string, *httpParcer.ErrorHTTP) {
	return "", nil
}

func (a *AuthService) CreateUser(c context.Context, userData models.RegistrationUser) (*models.User, *httpParcer.ErrorHTTP) {
	userData.Password = a.generateHash(userData.Password)

	data, err := a.repo.Store.CreateUser(c, userData)
	if err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Msg:  err.Error(),
			Code: 500,
		}
	}

	return data, nil
}

func (a *AuthService) generateHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte(solt)))
}
