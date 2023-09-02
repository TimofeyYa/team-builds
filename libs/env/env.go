package env

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

// Загрузка файла .env
func LoadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnln("Error loading .env file")
	}
}
