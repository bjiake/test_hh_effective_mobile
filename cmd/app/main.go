package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"hh.ru/cmd/internal/app"
	"hh.ru/cmd/internal/repo/car"
	"hh.ru/cmd/internal/repo/people"
	"log"
	"os"
	"time"
)

func main() {
	//Загрузка переменных окружения (если используете godotenv)
	err := godotenv.Load("app.env") // Раскомментируйте эту строку, если используете godotenv
	if err != nil {
		log.Fatal("Error loading .env file")
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
		log.Fatal(err)
	}
	defer db.Close()

	carRepository := car.NewPgSqlCarRepository(db)
	peopleRepository := people.NewPgSqlPeopleRepository(db)

	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	app.RunCarRepositoryDemo(ctx, *carRepository)
	app.RunPeopleRepositoryDemo(ctx, peopleRepository)
	//router.App()
}
