package db

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// ConnectToBD Подключение к PostgresSql по app.env
func ConnectToBD() (*sql.DB, error) {
	//Загрузка переменных окружения (если используете godotenv)
	err := godotenv.Load("app.env") // Раскомментируйте эту строку, если используете godotenv
	if err != nil {
		return nil, err
	}

	// Получение значений переменных окружения
	psqlUser := os.Getenv("PSQL_USER")
	psqlPass := os.Getenv("PSQL_PASS")
	psqlDBName := os.Getenv("PSQL_DBNAME")
	psqlPort := os.Getenv("PSQL_PORT")

	// Формирование строки подключения
	psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost:%s/%s", psqlUser, psqlPass, psqlPort, psqlDBName)

	// Подключение к БД
	db, err := sql.Open("pgx", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CloseDB Закрытие подключения
func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
