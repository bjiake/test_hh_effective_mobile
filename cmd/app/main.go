package main

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"hh.ru/pkg/config"
	"hh.ru/pkg/di"
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
	cfg, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}
	server, diErr := di.InitializeAPI(cfg)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
