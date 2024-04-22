package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	PsqlUser   string
	PsqlPass   string
	PsqlHost   string
	PsqlPort   string
	PsqlDBName string
}

func LoadConfig() (Config, error) {
	var config Config

	//Загрузка переменных окружения (если используете godotenv)
	err := godotenv.Load("app.env") // Раскомментируйте эту строку, если используете godotenv
	if err != nil {
		return config, err
	}
	// Получение значений переменных окружения
	config.PsqlUser = os.Getenv("PSQL_USER")
	config.PsqlPass = os.Getenv("PSQL_PASS")
	config.PsqlHost = os.Getenv("PSQL_HOST")
	config.PsqlDBName = os.Getenv("PSQL_DBNAME")
	config.PsqlPort = os.Getenv("PSQL_PORT")

	return config, err
}
