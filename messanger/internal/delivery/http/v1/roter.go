package v1

import (
	"teamBuild/messages/internal/service"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, s *service.Service) {
	v1 := router.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login")
			auth.POST("/registration")
			auth.POST("/authorization")
		}

		user := v1.Group("/user")
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
		}
	}
}
