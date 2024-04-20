package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
	"hh.ru/cmd/internal/app"
	"hh.ru/cmd/internal/repo"
	"log"
	"time"
)

func main() {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	carRepository := repo.NewPgSqlClassicRepository(db)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	app.RunRepositoryDemo(ctx, *carRepository)
	//router.App()
}
