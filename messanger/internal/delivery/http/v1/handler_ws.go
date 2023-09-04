package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type userMessage struct {
	Msg         string `json:"msg"`
	RecipientId int    `json:"recipient_id"`
}

var mu sync.Mutex // Объявляет мьютекс

func wshandler(w http.ResponseWriter, r *http.Request, userId int, userData map[int]*websocket.Conn) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Warnf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	mu.Lock()
	userData[userId] = conn
	mu.Unlock()

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(userData, userId)
			fmt.Println(userData)
			mu.Unlock()
			break
		}

		var userMsg userMessage
		if err := json.Unmarshal(msg, &userMsg); err != nil {
			conn.WriteMessage(t, []byte(err.Error()))
			continue
		}

		mu.Lock()
		if userData[userMsg.RecipientId] != nil {
			userData[userMsg.RecipientId].WriteMessage(t, msg)
		}
		mu.Unlock()
	}
	conn.Close()
}

type messageRequest struct {
	UserId int `uri:"user_id"`
}

func (h *Handler) MessageChanal(c *gin.Context, userData map[int]*websocket.Conn) {
	var messReq messageRequest
	if err := c.ShouldBindUri(&messReq); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"status": false, "message": "user id is required"})
		return
	}
	wshandler(c.Writer, c.Request, messReq.UserId, userData)
}
