package app

import (
	"fmt"
	"os"
	"teamBuild/messages/internal/delivery/http"
	"teamBuild/messages/internal/repository"
	"teamBuild/messages/internal/service"
	clinets "teamBuilds/libs/clients"
	"teamBuilds/libs/env"
	"teamBuilds/libs/logger"

	"github.com/sirupsen/logrus"
)

func Run() {
	// Configuration
	logger.InitLogger()
	env.LoadEnvFile()

	// HTTP Server configurate
	cnf := clinets.DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DataBase: os.Getenv("DB_NAME"),
	}
	fmt.Println(cnf)
	pgPool, err := clinets.InitPostgresClient(cnf, 5)
	if err != nil {
		logrus.Fatalf("Error connect to data base: %s", err.Error())
	}

	repo := repository.NewRepository(pgPool)
	service := service.NewService(repo)
	handler := http.NewHandler(service)

	httpServer := http.CreateHTTPServer("8080", handler.InitRoutes())

	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("Error start server: %s", err.Error())
	}
}
