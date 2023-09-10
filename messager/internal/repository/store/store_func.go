package store

import (
	"context"
	"fmt"
	"teamBuild/messages/internal/models"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func (s *StoreRepository) CreateUser(c context.Context, userData models.RegistrationUser) (*models.User, error) {
	tx, err := s.pool.BeginTx(c, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}
	defer s.rollBackTx(tx, c)

	sqlReq, args, err := psql.Insert("users").
		Columns("name", "email", "password_hash").
		Values(userData.Name, userData.Email, userData.Password).
		Suffix("RETURNING \"id\"").ToSql()

	if err != nil {
		logrus.Error(err.Error())
		return nil, fmt.Errorf("error build query: %s", err.Error())
	}

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

func (s *StoreRepository) LoginUser(c context.Context, cred models.Credentials) (int, error) { // Вернём id пользователя
	sqlReq, args, err := psql.Select("id").
		From("users").
		Where(sq.And{
			sq.Eq{
				"email": cred.Email,
			},
			sq.Eq{
				"password_hash": cred.Password,
			},
		}).ToSql()

	if err != nil {
		logrus.Error(err.Error())
		return 0, fmt.Errorf("error build query: %s", err.Error())
	}

	var userId int
	if err := s.pool.QueryRow(c, sqlReq, args...).Scan(&userId); err != nil {
		logrus.Error(err.Error())
		return 0, fmt.Errorf("error login: %s", err.Error())
	}

	return userId, nil
}

func (s *StoreRepository) SaveRefreshToken(c context.Context, userId int, refreshToken string) error {
	sqlReq, args, err := psql.Insert("sessions").
		Columns("user_id", "refresh_token", "expires_at", "created_at").
		Values(userId, refreshToken, time.Now().Add(48*time.Hour), time.Now()).ToSql()
	if err != nil {
		logrus.Error(err.Error())
		return fmt.Errorf("error build query: %s", err.Error())
	}

	_, err = s.pool.Exec(c, sqlReq, args...)
	if err != nil {
		logrus.Error(err.Error())
		return fmt.Errorf("error exec query: %s", err.Error())
	}

	return nil
}

func (s *StoreRepository) UpdateRefreshToken(c context.Context, userId int, oldRefresh string, newRefresh string) error {
	tx, err := s.pool.BeginTx(c, pgx.TxOptions{
		IsoLevel: pgx.ReadCommitted,
	})
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	defer s.rollBackTx(tx, c)

	sqlReq, args, err := psql.Update("sessions").
		Set("refresh_token", newRefresh).
		Set("expires_at", time.Now().Add(48*time.Hour)).
		Where(sq.And{
			sq.Eq{
				"refresh_token": oldRefresh,
				"user_id":       userId,
			},
			sq.Gt{
				"expires_at": time.Now(),
			},
		}).ToSql()
	if err != nil {
		logrus.Error(err.Error())
		return fmt.Errorf("error build query: %s", err.Error())
	}

	result, err := s.pool.Exec(c, sqlReq, args...)
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	if result.RowsAffected() != 1 {
		return fmt.Errorf("refresh token in not valid")
	}

	return nil
}

func (s *StoreRepository) CreateMessage(c context.Context, msg *models.Message) (int, error) {
	sqlReq, args, err := psql.Insert("messages").
		Columns("sendler_id", "recipient_id", "message", "is_read", "created_at").
		Values(msg.SenderId, msg.RecipientId, msg.Content, msg.IsRead, msg.CreateAt).
		Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		logrus.Error(err.Error())
		return 0, fmt.Errorf("error build query: %s", err.Error())
	}

	var messageId int
	if err := s.pool.QueryRow(c, sqlReq, args...).Scan(&messageId); err != nil {
		if err != nil {
			logrus.Error(err.Error())
			return 0, fmt.Errorf("error exec query: %s", err.Error())
		}
	}

	return messageId, nil
}
