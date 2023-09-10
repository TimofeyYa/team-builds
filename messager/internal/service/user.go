package service

import (
	"context"
	"teamBuild/messages/internal/models"
	"teamBuild/messages/internal/repository"
	httpParcer "teamBuilds/libs/http_parcer"
	"time"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
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
func (u *UserService) CreateMessage(c context.Context, userId int, recipientId int, msg string) (*models.Message, *httpParcer.ErrorHTTP) {
	msgItem := models.Message{
		Content:     msg,
		SenderId:    userId,
		RecipientId: recipientId,
		IsRead:      false,
		CreateAt:    time.Now(),
	}
	messageId, err := u.repo.CreateMessage(c, &msgItem)
	if err != nil {
		return nil, &httpParcer.ErrorHTTP{
			Code: 500,
			Msg:  "Create message err",
		}
	}
	msgItem.Id = messageId
	return &msgItem, nil
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
