package v1

import (
	"teamBuild/messages/internal/service"
	httpparcer "teamBuilds/libs/http_parcer"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SetRouter(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", httpparcer.Parce(h.Login))
			auth.POST("/registration", httpparcer.Parce(h.Registration))
			auth.PUT("/authorization", h.Authorization)
		}

		user := v1.Group("/user", h.validateJWT)
		{
			// Get information about user
			user.GET("/:user_id")

			// Get information about users friends
			user.GET("/:user_id/fiends")

			// Create friend for user (only for user)
			user.POST("/:user_id/fiends")

			// Delete friend for user (only for user)
			user.DELETE("/:user_id/fiends")

			// Get user chats (only for user)
			user.GET("/:user_id/chats")

			// Web Socket Chanal for messages (only for user)
			user.GET("/:user_id/chats/ws", CreateChatConnection(h.MessageChanal))

			// Get user chat with recipient  (only for user)
			user.GET("/:user_id/chats/:recipient_id")

			// Delete user chat with recipient  (only for user)
			user.DELETE("/:user_id/chats/:recipient_id")

			// Create message for chat  (only for user)
			user.POST("/:user_id/chats/:recipient_id")

			// Update message for chat  (only for user)
			user.PATCH("/:user_id/chats/:recipient_id/:message_id")
		}
	}
}

func CreateChatConnection(f func(*gin.Context, map[int]*websocket.Conn)) func(*gin.Context) {
	userData := map[int]*websocket.Conn{}
	return func(c *gin.Context) {
		f(c, userData)
	}
}

func (h *Handler) errorResponse(c *gin.Context, code uint16, errMsg string) {
	c.AbortWithStatusJSON(int(code), gin.H{
		"status": false,
		"error":  errMsg,
	})
}
