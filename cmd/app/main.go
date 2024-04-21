package main

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	//db := repo.ConnectToBD()
	//defer repo.CloseDB(db)

	//ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	//defer cancel()

	//carRepository := car.NewPgSqlCarRepository(db)
	//app.RunCarRepositoryDemo(context.Background(), *carRepository)
	//peopleRepository := people.NewPgSqlPeopleRepository(db)
	//app.RunPeopleRepositoryDemo(context.Background(), peopleRepository)
	//router.App(db)
	if err := run(context.Background()); err != nil {
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	server, err := rest.NewServer()
	if err != nil {
		return err
	}
	err = server.Run(ctx)
	return err
}
