package di

import (
	http "hh.ru/pkg/api"
	"hh.ru/pkg/api/handler"
	"hh.ru/pkg/config"
	"hh.ru/pkg/db"
	"hh.ru/pkg/repo/car"
	"hh.ru/pkg/repo/people"
	"hh.ru/pkg/service"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	bd, err := db.ConnectToBD(cfg)
	if err != nil {
		return nil, err
	}
	carRepository := car.NewCarRepository(bd)
	peopleRepository := people.NewPeopleRepository(bd)
	userService := service.NewService(carRepository, peopleRepository)
	userHandler := handler.NewHandler(userService)
	serverHTTP := http.NewServerHTTP(userHandler)
	return serverHTTP, nil
}
