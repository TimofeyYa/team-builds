package repository

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository/store"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store interface {
	CreateUser(context.Context, models.RegistrationUser) (*models.User, error)
}

type Repository struct {
	Store
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Store: store.NewStoreRetository(pool),
	}
}
