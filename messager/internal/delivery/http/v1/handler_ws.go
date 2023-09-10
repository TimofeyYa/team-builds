package v1

import (
	"context"
	"encoding/json"
	"sync"
	"teamBuild/messages/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// TODO: Добавить время отправления сразу в сообщение
type userMessage struct {
	MsgType     string `json:"msg_type" binding:"required"`
	RecipientId int    `json:"recipient_id" binding:"required"`
	Msg         string `json:"msg"`    // Нужен для создания и обновления
	MsgId       int    `json:"msg_id"` // Нужен для удаления и обновления
}

var mu sync.Mutex // Объявляет мьютекс

type WSHandler struct {
	service         *service.Service
	userConnections map[int]*websocket.Conn
}

func NewWSHandler(service *service.Service) *WSHandler {
	userConnections := map[int]*websocket.Conn{}
	return &WSHandler{
		userConnections: userConnections,
		service:         service,
	}
}

func (wsh *WSHandler) wshandler(c context.Context, conn *websocket.Conn, userId int) {
	mu.Lock()
	wsh.userConnections[userId] = conn
	mu.Unlock()
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(wsh.userConnections, userId)
			mu.Unlock()
			break
		}

		var userMsg userMessage
		if err := json.Unmarshal(msg, &userMsg); err != nil {
			conn.WriteMessage(t, []byte(err.Error()))
			continue
		}

		switch userMsg.MsgType {
		case "read":
			{
				wsh.service.ReadMessages(c, userId, userMsg.RecipientId)
			}
		case "create":
			{
				wsh.service.CreateMessage(c, userId, userMsg.RecipientId, userMsg.Msg)
			}
		case "delete":
		case "update":
		default:
			{
				conn.WriteMessage(t, []byte("Не валидный тип сообщения"))
				continue
			}
		}

		mu.Lock()
		if wsh.userConnections[userMsg.RecipientId] != nil {
			wsh.userConnections[userMsg.RecipientId].WriteMessage(t, msg)
		}
		mu.Unlock()
	}
}

type messageRequest struct {
	UserId int `uri:"user_id"`
}

func (wsh *WSHandler) MessageChanal(c *gin.Context) {
	var messReq messageRequest
	if err := c.ShouldBindUri(&messReq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "message": "user id is required"})
		return
	}

	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Warnf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	wsh.wshandler(c, conn, messReq.UserId)

	if err := conn.Close(); err != nil {
		logrus.Errorf("Failed to close websocket connect: %+v\n", err)
	}
}
