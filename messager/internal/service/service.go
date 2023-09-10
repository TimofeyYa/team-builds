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
	GetUserInfo(context.Context, int) (*models.User, *httpParcer.ErrorHTTP)
	GetUserFriends(c context.Context, userId int, limit int, offset int) ([]models.User, *httpParcer.ErrorHTTP)
	CreateUserFriend(c context.Context, userId int, friendId int) *httpParcer.ErrorHTTP
	DeleteUserFriend(c context.Context, userId int, friendId int) *httpParcer.ErrorHTTP
	GetUserChats(c context.Context, userId int) ([]models.Chat, *httpParcer.ErrorHTTP)
	CreateMessage(c context.Context, userId int, recipientId int, msg string) (*models.Message, *httpParcer.ErrorHTTP)
	UpdateMessage(c context.Context, userId int, messageId int, msg models.Message) *httpParcer.ErrorHTTP
	DeleteMessage(c context.Context, userId int, messageId int) *httpParcer.ErrorHTTP
	ReadMessages(c context.Context, userId int, recipientId int) *httpParcer.ErrorHTTP
	GetChatMessages(c context.Context, userId int, recipientId int) ([]models.Message, *httpParcer.ErrorHTTP)
}

type Service struct {
	Auth
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
		User: NewUserService(repo),
	}
}
