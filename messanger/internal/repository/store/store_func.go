package store

import (
	"context"
	"fmt"
	"teamBuild/messages/internal/models"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (s *StoreRepository) CreateUser(c context.Context, userData models.RegistrationUser) (*models.User, error) {
	sqlReq, args, err := psql.Insert("users").
		Columns("name", "email", "password_hash").
		Values(userData.Name, userData.Email, userData.Password).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		logrus.Error(err.Error())
		return nil, fmt.Errorf("error build query: %s", err.Error())
	}

	tx, err := s.pool.BeginTx(c, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	defer s.rollBackTx(tx, c)
	var userId int
	if err := tx.QueryRow(c, sqlReq, args...).Scan(&userId); err != nil {
		return nil, fmt.Errorf("error query: %s", err.Error())
	}

	getUserReq, args, err := psql.Select("*").
		From("users").Where(sq.Eq{
		"id": userId,
	}).ToSql()
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	var newUserData models.User
	if err := tx.QueryRow(c, getUserReq, args...).Scan(
		&newUserData.Id,
		&newUserData.Name,
		&newUserData.Email,
		&newUserData.Password,
		&newUserData.UpdateAt,
		&newUserData.CreateAt,
	); err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	if err := tx.Commit(c); err != nil {
		return nil, fmt.Errorf("error commit: %s", err.Error())
	}

	return &newUserData, nil
}
