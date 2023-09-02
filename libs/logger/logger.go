package logger

import "github.com/sirupsen/logrus"

// Конфигурирация логгера
func InitLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
}
