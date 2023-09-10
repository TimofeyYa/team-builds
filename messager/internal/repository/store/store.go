package store

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type StoreRepository struct {
	pool *pgxpool.Pool
}

func NewStoreRetository(pool *pgxpool.Pool) *StoreRepository {
	return &StoreRepository{
		pool: pool,
	}
}

func (s *StoreRepository) rollBackTx(tx pgx.Tx, ctx context.Context) {
	err := tx.Rollback(ctx)
	if err != nil {
		if !errors.Is(err, pgx.ErrTxClosed) {
			log.Println(err.Error())
		}
	}
}
