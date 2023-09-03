package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type userMessage struct {
	Msg         string `json:"msg"`
	RecipientId int    `json:"recipient_id"`
}

func wshandler(w http.ResponseWriter, r *http.Request, userId int, userData map[int]*websocket.Conn) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	userData[userId] = conn
	fmt.Println(userData)
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			delete(userData, userId)
			break
		}

		var userMsg userMessage
		if err := json.Unmarshal(msg, &userMsg); err != nil {
			conn.WriteMessage(t, []byte(err.Error()))
			continue
		}

		if userData[userMsg.RecipientId] != nil {
			userData[userMsg.RecipientId].WriteMessage(t, msg)
		}
		conn.WriteMessage(t, msg)
	}
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
