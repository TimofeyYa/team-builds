package repository

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository/store"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	CreateUser(context.Context, models.RegistrationUser) (*models.User, error)
	LoginUser(context.Context, models.Credentials) (int, error)
	SaveRefreshToken(context.Context, int, string) error
	UpdateRefreshToken(context.Context, int, string, string) error
}

type Repository struct {
	Store
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Store: store.NewStoreRetository(pool),
	}
}
