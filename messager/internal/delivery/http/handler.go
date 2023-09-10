package http

import (
	v1 "teamBuild/messages/internal/delivery/http/v1"
	"teamBuild/messages/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	handlerV1 := v1.NewHandler(h.service)

	h.setCors(router)
	handlerV1.SetRouter(router)

	return router
}

// Задаём настройки корс
func (h *Handler) setCors(e *gin.Engine) {
	e.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},

		AllowCredentials: true,
	}))

	e.Use(CORSMiddleware())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
