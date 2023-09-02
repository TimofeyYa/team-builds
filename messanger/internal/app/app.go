package app

import (
	"teamBuild/messages/internal/delivery/http"
	"teamBuilds/libs/logger"

	"github.com/sirupsen/logrus"
)

func Run() {
	logger.InitLogger()

	handler := http.NewHandler()

	httpServer := http.CreateHTTPServer("8080", handler.InitRoutes())

	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("Error start server: %s", err.Error())
	}
}
