package app

import (
	"teamBuild/messages/internal/delivery/http"
	"teamBuild/messages/internal/service"
	"teamBuilds/libs/env"
	"teamBuilds/libs/logger"

	"github.com/sirupsen/logrus"
)

func Run() {
	// Configuration
	logger.InitLogger()
	env.LoadEnvFile()

	// Server start
	service := service.NewService()
	handler := http.NewHandler(service)

	httpServer := http.CreateHTTPServer("8080", handler.InitRoutes())

	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("Error start server: %s", err.Error())
	}
}
