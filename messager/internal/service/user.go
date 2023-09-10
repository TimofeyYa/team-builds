package service

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (u *UserService) GetUserInfo(context.Context, int) (*models.User, *httpParcer.ErrorHTTP) {
	return nil, nil
}
func (u *UserService) GetUserFriends(c context.Context, userId int, limit int, offset int) ([]models.User, *httpParcer.ErrorHTTP) {
	return nil, nil
}
func (u *UserService) CreateUserFriend(c context.Context, userId int, friendId int) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) DeleteUserFriend(c context.Context, userId int, friendId int) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) GetUserChats(c context.Context, userId int) ([]models.Chat, *httpParcer.ErrorHTTP) {
	return nil, nil
}
func (u *UserService) CreateMessage(c context.Context, userId int, recipientId int, msg models.Message) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) UpdateMessage(c context.Context, userId int, messageId int, msg models.Message) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) DeleteMessage(c context.Context, userId int, messageId int) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) ReadMessages(c context.Context, userId int, recipientId int) *httpParcer.ErrorHTTP {
	return nil
}
func (u *UserService) GetChatMessages(c context.Context, userId int, recipientId int) ([]models.Message, *httpParcer.ErrorHTTP) {
	return nil, nil
}
