package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const solt = "message"

type AuthService struct {
	secretKey string
	repo      *repository.Repository
}

func NewAuthService(repo *repository.Repository) *AuthService {
	return &AuthService{
		secretKey: os.Getenv("APP_SECRET_KEY"),
		repo:      repo,
	}
}

func (a *AuthService) LoginUser(c context.Context, cred models.Credentials) (string, string, *httpParcer.ErrorHTTP) {
	cred.Password = a.generateHash(cred.Password)

	userId, err := a.repo.Store.LoginUser(c, cred)
	if err != nil {
		return "", "", &httpParcer.ErrorHTTP{
			Code: 401,
			Msg:  "Не правильный логин или пароль",
		}
	}

	token, err := a.createJWT(userId)
	if err != nil {
		logrus.Errorf("err JWT: %s", err.Error())
		return "", "", &httpParcer.ErrorHTTP{
			Code: 500,
			Msg:  "Внутренняя ошибка сервера",
		}
	}
	refreshToken := uuid.NewString()
	err = a.repo.SaveRefreshToken(c, userId, refreshToken)
	if err != nil {
		logrus.Errorf("err Refresh token: %s", err.Error())
		return "", "", &httpParcer.ErrorHTTP{
			Code: 500,
			Msg:  "Внутренняя ошибка сервера",
		}
	}

	return token, refreshToken, nil
}

func (a *AuthService) CreateUser(c context.Context, userData models.RegistrationUser) (*models.User, *httpParcer.ErrorHTTP) {
	// TODO: Добавить валидацию имени, пароля и почты
	userData.Password = a.generateHash(userData.Password)

	data, err := a.repo.Store.CreateUser(c, userData)

	// TODO: Добавить проверку ошибок и в зависимости от типа ошибки выдавать код
	if err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Msg:  err.Error(),
			Code: 500,
		}
	}

	return data, nil
}

func (a *AuthService) Authorization(c context.Context, tokens *models.TokenPair) (*models.TokenPair, *httpParcer.ErrorHTTP) {
	tokenContent, err := a.parseJWT(tokens.JWT)
	// TODO: Мы должны различать, когда токен не валидный из за времени, а когда из за структуры к ключа
	// Если токен не валиден из за времени, то мы заменяем его рефрешем, если из за ключа - отказываем в доступе
	if err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Msg:  err.Error(),
			Code: 403,
		}
	}

	newRefresh := uuid.NewString()
	if err := a.repo.Store.UpdateRefreshToken(c, tokenContent.UserId, tokens.RefreshToken, newRefresh); err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Msg:  err.Error(),
			Code: 403,
		}
	}

	newJwt, err := a.createJWT(tokenContent.UserId)
	if err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Msg:  err.Error(),
			Code: 500,
		}
	}

	return &models.TokenPair{
		JWT:          newJwt,
		RefreshToken: newRefresh,
	}, nil
}

func (a *AuthService) generateHash(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte(solt)))
}

type JWTContent struct {
	jwt.RegisteredClaims
	UserId int
}

func (a *AuthService) createJWT(userId int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTContent{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(2 * time.Hour)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		UserId: userId,
	})

	return token.SignedString([]byte(a.secretKey))
}

func (a *AuthService) parseJWT(accessToken string) (*JWTContent, error) {
	fmt.Println(accessToken)
	token, err := jwt.ParseWithClaims(accessToken, &JWTContent{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(a.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	tokenData, ok := token.Claims.(*JWTContent)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return tokenData, err
}
