package app

import (
	"log"
	"teamBuild/messages/internal/delivery/http"
)

func Run() {
	handler := http.NewHandler()

	httpServer := http.CreateHTTPServer("8080", handler.InitRoutes())

	if err := httpServer.Run(); err != nil {
		log.Println("321")
	}
}
